package page_func

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type response_data struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious int `json:"malicious"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

type return_func_data_bac struct {
	HASH   string `json:"hash"`
	Namber int    `json:"namber"`
	Status int    `json:"status"`
}

func calculateFileHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func uploadFileToVirusTotal(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Не вдалося відкрити файл для завантаження:", err)
		return ""
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filePath)

	io.Copy(part, file)

	writer.Close()

	req, err := http.NewRequest("POST", "https://www.virustotal.com/api/v3/files", body)
	if err != nil {
		fmt.Println("Помилка створення запиту:", err)
		return ""
	}

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Помилка запиту:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Помилка при завантаженні файлу: %d\n", resp.StatusCode)
		return ""
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fileID := result["data"].(map[string]interface{})["id"].(string)

	return fileID
}

func checkFileSecurityStatus(fileID string) (int, int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.virustotal.com/api/v3/files/"+fileID, nil)
	if err != nil {
		fmt.Println("Помилка створення запиту:", err)
		return 0, 0
	}

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Помилка запиту:", err)
		return 0, 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Помилка при отриманні результату аналізу: %d\n", resp.StatusCode)
		return 1, 0
	}

	var result response_data
	json.NewDecoder(resp.Body).Decode(&result)

	maliciousCount := result.Data.Attributes.LastAnalysisStats.Malicious

	return 2, maliciousCount
}

func Scan_file_virus(nameFilePath string) return_func_data_bac {
	var return_func return_func_data_bac
	hash := calculateFileHash(nameFilePath)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", config_main.Files_virustotal+hash, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Помилка запиту:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fileID := uploadFileToVirusTotal(nameFilePath)

		if fileID != "" {
			data, data_1 := checkFileSecurityStatus(fileID)

			if data == 1 {
				return_func.Status = 1
			} else if data == 2 {
				return_func.HASH = hash
				return_func.Namber = data_1
				return_func.Status = 2
			}
		}
	} else if resp.StatusCode == http.StatusOK {
		data, data_1 := checkFileSecurityStatus(hash)

		if data == 1 {
			return_func.Status = 1
		} else if data == 2 {
			return_func.HASH = hash
			return_func.Namber = data_1
			return_func.Status = 2
		}
	}

	return return_func
}

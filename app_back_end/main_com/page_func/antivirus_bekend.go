package page_func

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	Data   string `json:"data"`
}

func calculateFileHash(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

func uploadFileToVirusTotal(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filePath)

	io.Copy(part, file)

	writer.Close()

	req, _ := http.NewRequest("POST", "https://www.virustotal.com/api/v3/files", body)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: %d\n", resp.StatusCode)
		return ""
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fileID := result["data"].(map[string]interface{})["id"].(string)

	return fileID
}

func checkFileSecurityStatus(fileID string) (int, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.virustotal.com/api/v3/files/"+fileID, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: %d\n", resp.StatusCode)
		return 1, 0
	}

	var result response_data
	json.NewDecoder(resp.Body).Decode(&result)

	maliciousCount := result.Data.Attributes.LastAnalysisStats.Malicious

	return 2, maliciousCount
}

func start_data_exe(file string) string {
	exePath := config_main.Antivirus_data_exe
	workingDir := config_main.Library_folder
	dataFilePath := "./" + config_main.Library_folder + "/data/" + config_main.File_data_exe

	cmd := exec.Command(exePath, "../"+file)
	cmd.Dir = workingDir

	cmd.Start()
	cmd.Wait()

	data, _ := ioutil.ReadFile(dataFilePath)
	func_all.Clear_file(config_main.Global_phat + "\\" + config_main.Library_folder + "\\data\\" + config_main.File_data_exe)

	return string(data)
}

func Scan_file_virus(nameFilePath string) return_func_data_bac {
	var return_func return_func_data_bac
	hash := calculateFileHash(nameFilePath)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", config_main.Files_virustotal+hash, nil)

	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	data_file := start_data_exe(nameFilePath)

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
				return_func.Data = data_file
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
			return_func.Data = data_file
		}
	}

	return return_func
}

// scan dir// scan dir// scan dir// scan dir// scan dir
// scan dir// scan dir// scan dir// scan dir// scan dir
// scan dir// scan dir// scan dir// scan dir// scan dir

func File_hash(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil))
}

func Check_hash_VirusTotal(fileHash string) map[string]interface{} {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", fileHash)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-apikey", config_main.ApiKey_virustotal)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return map[string]interface{}{"message": "Hash not found in database VirusTotal"}
	} else if resp.StatusCode != http.StatusOK {
		return nil
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

func Scan_exeFiles(rootDir string) []string {
	var exeFiles []string
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".exe") {
			exeFiles = append(exeFiles, path)
		}
		return nil
	})
	return exeFiles
}

func Delete_file(filePath string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")

	if err := os.Remove(filePath); err != nil {
		return "0"
	}

	return "1"
}

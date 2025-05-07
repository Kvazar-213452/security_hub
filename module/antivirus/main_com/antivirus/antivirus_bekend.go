package antivirus

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io"
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
	Data   string `json:"data"`
}

func calculateFileHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file:", err)
		return ""
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		fmt.Println("error hashing file:", err)
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func checkFileByHash(hash string) (int, int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.virustotal.com/api/v3/files/"+hash, nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return 1, 0
	}
	req.Header.Add("x-apikey", config_main.ApiKey_virustotal)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request:", err)
		return 1, 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 3, 0
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: %d\n", resp.StatusCode)
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
	if hash == "" {
		return_func.Status = 1
		return return_func
	}

	status, malicious := checkFileByHash(hash)
	if status == 2 {
		return_func.HASH = hash
		return_func.Namber = malicious
		return_func.Status = 2
		return_func.Data = "0"
	} else if status == 3 {
		return_func.HASH = hash
		return_func.Namber = malicious
		return_func.Status = 2
		return_func.Data = "1"
	} else {
		return_func.HASH = hash
		return_func.Namber = malicious
		return_func.Status = 2
		return_func.Data = "2"
	}

	return return_func
}

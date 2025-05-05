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
	"path/filepath"
	"strings"
)

// app_back_end/main_com/page_func/antivirus/scan_dir.go

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
		if err != nil {
			return nil
		}

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

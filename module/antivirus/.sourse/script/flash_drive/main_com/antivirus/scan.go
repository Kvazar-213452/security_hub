package antivirus

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const VirusTotalAPIKey = "b022e26c03533fcd236535e650661b72c41165db486500bf5877f08184b21099"

type ScanResult struct {
	AllExeFiles    []string `json:"all_exe_files"`
	MaliciousFiles []string `json:"malicious_files"`
}

func ReadFlashDriveLetter(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	path := strings.TrimSpace(string(content))
	if !strings.HasSuffix(path, "\\") {
		path += "\\"
	}
	return path, nil
}

func FindExeFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".exe") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func CalculateSHA256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func CheckVirusTotalHash(hash string) (bool, error) {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s", hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("x-apikey", VirusTotalAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return false, nil
	}
	if resp.StatusCode != 200 {
		return false, fmt.Errorf("error from VirusTotal: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return false, nil
	}

	attributes, ok := data["attributes"].(map[string]interface{})
	if !ok {
		return false, nil
	}

	stats, ok := attributes["last_analysis_stats"].(map[string]interface{})
	if !ok {
		return false, nil
	}

	if malicious, ok := stats["malicious"].(float64); ok && malicious > 0 {
		return true, nil
	}
	return false, nil
}

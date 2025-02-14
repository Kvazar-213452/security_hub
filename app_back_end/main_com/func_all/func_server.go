package func_all

import (
	"encoding/json"
	config_main "head/main_com/config"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// app_back_end/main_com/func_all/func_server.go

type LogContent struct {
	Log string `json:"log"`
}

func LoadLogFile() []byte {
	content, _ := ioutil.ReadFile(config_main.Log_file)

	logContent := LogContent{
		Log: string(content),
	}

	jsonData, _ := json.Marshal(logContent)

	return jsonData
}

func RemoveNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func splitLines(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n")
}

func ClearDirectory(dir string) error {
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		err := os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func Check_server_Status(url string) int {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return 1
	}

	return 0
}

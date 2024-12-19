package func_all

import (
	"encoding/json"
	config_main "head/main_com/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

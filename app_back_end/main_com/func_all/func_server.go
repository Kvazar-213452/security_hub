package func_all

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type LogContent struct {
	Log string `json:"log"`
}

func LoadLogFile() ([]byte, error) {
	content, err := ioutil.ReadFile(config_main.Log_file)
	if err != nil {
		return nil, fmt.Errorf("не вдалося прочитати файл: %w", err)
	}

	logContent := LogContent{
		Log: string(content),
	}

	jsonData, err := json.Marshal(logContent)
	if err != nil {
		return nil, fmt.Errorf("не вдалося закодувати в JSON: %w", err)
	}

	return jsonData, nil
}

func RemoveNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func splitLines(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n")
}

func ClearDirectory(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

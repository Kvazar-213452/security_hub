package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Get_phat_global() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return ""
	}

	exeDir := filepath.Dir(exePath)
	return exeDir
}

func Get_server_url(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return ""
	}

	server, ok := data["server"].(string)
	if !ok {
		return ""
	}

	return server
}

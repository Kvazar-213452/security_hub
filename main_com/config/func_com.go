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

	server, ok := data["url"].(string)
	if !ok {
		return ""
	}

	return server
}

func Get_antivirus_flash_drive_cmd(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	var config Config_global
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return ""
	}

	antivirus_flash_drive_cmd := config.Antivirus.Antivirus_flash_drive_cmd
	if antivirus_flash_drive_cmd == "" {
		return ""
	}

	return antivirus_flash_drive_cmd
}

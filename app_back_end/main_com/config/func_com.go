package config

// app_back_end/main_com/config/func_com.go

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Get_phat_global() string {
	exePath, _ := os.Executable()

	exeDir := filepath.Dir(exePath)
	return exeDir
}

func Get_antivirus_flash_drive_cmd(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	var config Config_global
	json.NewDecoder(file).Decode(&config)

	antivirus_flash_drive_cmd := config.Antivirus.Antivirus_flash_drive_cmd
	if antivirus_flash_drive_cmd == "" {
		return ""
	}

	return antivirus_flash_drive_cmd
}

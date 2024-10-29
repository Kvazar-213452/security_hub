package config

import (
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

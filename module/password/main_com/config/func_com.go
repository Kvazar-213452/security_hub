package config

// module/password/main_com/config/func_com.go

import (
	"os"
	"path/filepath"
)

func Get_phat_global() string {
	exePath, _ := os.Executable()

	exeDir := filepath.Dir(exePath)
	return exeDir
}

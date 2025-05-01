package config

import (
	"os"
	"path/filepath"
)

// module/system/main_com/config/func_com.go

func Get_phat_global() string {
	exePath, _ := os.Executable()

	exeDir := filepath.Dir(exePath)
	return exeDir
}

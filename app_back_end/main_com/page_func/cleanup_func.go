package page_func

import (
	"fmt"
	config_main "head/main_com/config"
	"os/exec"
	"syscall"
)

func Cleanup() {
	exePath := "cleanup_script.bat"
	workingDir := "library"

	cmd := exec.Command("powershell.exe", "Start-Process", exePath, "-WorkingDirectory", workingDir, "-Verb", "runAs")

	cmd.Run()

	cleanupDLL, err := syscall.LoadDLL(config_main.Cleanup_dll)
	if err != nil {
		fmt.Printf("Не вдалося завантажити DLL: %v\n", err)
		return
	}
	defer cleanupDLL.Release()

	cleanupProc, err := cleanupDLL.FindProc("cleanup")
	if err != nil {
		return
	}

	cleanupProc.Call()
}

func Cleanup_wifi() {
	exePath := "wifi.bat"
	workingDir := "library/cleanup"

	cmd := exec.Command("cmd", "/C", "start", "/min", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_backup() {
	exePath := "backup.bat"
	workingDir := "library/cleanup"

	cmd := exec.Command("cmd", "/C", "start", "/min", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_desktop() {
	exePath := "desktop.bat"
	workingDir := "library/cleanup"

	cmd := exec.Command("cmd", "/C", "start", "/min", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_doskey() {
	exePath := "doskey.bat"
	workingDir := "library/cleanup"

	cmd := exec.Command("cmd", "/C", "start", "/min", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

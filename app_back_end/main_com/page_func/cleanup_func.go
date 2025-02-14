package page_func

import (
	config_main "head/main_com/config"
	"os/exec"
	"syscall"
)

// app_back_end/main_com/page_func/cleanup_func.go

func Cleanup() {
	exePath := "cleanup_script.bat"
	workingDir := "library"

	cmd := exec.Command("powershell.exe", "Start-Process", exePath, "-WorkingDirectory", workingDir, "-Verb", "runAs")

	cmd.Run()

	cleanupDLL, _ := syscall.LoadDLL(config_main.Cleanup_dll)
	defer cleanupDLL.Release()

	cleanupProc, _ := cleanupDLL.FindProc("cleanup")
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

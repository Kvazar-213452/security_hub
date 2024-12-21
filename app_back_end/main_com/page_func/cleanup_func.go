package page_func

import (
	"os/exec"
)

func Cleanup() {
	cleanup()
}

func runDiskCleanup() error {
	cmd := exec.Command("sudo", "apt-get", "clean")
	cmd.Run()

	cmd = exec.Command("sudo", "apt-get", "autoremove", "--purge", "-y")
	cmd.Run()

	cmd = exec.Command("sudo", "rm", "-rf", "/tmp/*")
	cmd.Run()

	return nil
}

func emptyRecycleBin() int {
	cmd := exec.Command("gio", "trash", "--empty")
	cmd.Run()

	return 1
}

func cleanup() {
	emptyRecycleBin()
	runDiskCleanup()
}

func Cleanup_wifi() {
	exePath := "wifi.sh"
	workingDir := "library/cleanup"

	cmd := exec.Command("bash", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_backup() {
	exePath := "backup.sh"
	workingDir := "library/cleanup"

	cmd := exec.Command("bash", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_desktop() {
	exePath := "desktop.sh"
	workingDir := "library/cleanup"

	cmd := exec.Command("bash", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

func Cleanup_doskey() {
	exePath := "doskey.sh"
	workingDir := "library/cleanup"

	cmd := exec.Command("bash", exePath)
	cmd.Dir = workingDir

	cmd.Run()
}

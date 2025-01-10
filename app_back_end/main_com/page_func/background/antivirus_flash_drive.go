package background

import (
	"fmt"
	config_main "head/main_com/config"
	"os/exec"
	"strings"
	"time"
)

func Stop_antivirus_flash_drive_func() {
	config_main.Stop_antivirus_flash_drive <- true
}

func MonitorFlashDrives(stopChannel chan bool, command_cmd string) {
	var lastDrives []string

	for {
		select {
		case <-stopChannel:
			fmt.Println("Завершення моніторингу флешок.")
			return
		default:
			drives, err := getFlashDrives()
			if err != nil {
				fmt.Println("Помилка при отриманні флешок:", err)
				return
			}

			if len(drives) != len(lastDrives) || !equal(drives, lastDrives) {
				for _, drive := range drives {
					if !contains(lastDrives, drive) {
						fmt.Println("Підключено нову флешку:", drive)

						cmd := exec.Command("cmd", "/C", command_cmd)

						cmd.Start()
					}
				}
				lastDrives = drives
			}

			time.Sleep(2 * time.Second)
		}
	}
}

func getFlashDrives() ([]string, error) {
	cmd := exec.Command("wmic", "logicaldisk", "where", "DriveType=2", "get", "DeviceID,VolumeName")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	result := string(output)
	lines := strings.Split(result, "\n")

	var drives []string

	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line != "" {
			drives = append(drives, line)
		}
	}
	return drives, nil
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(arr []string, elem string) bool {
	for _, a := range arr {
		if a == elem {
			return true
		}
	}
	return false
}

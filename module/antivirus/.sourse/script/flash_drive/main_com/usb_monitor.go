package main_com

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type DriveInfo struct {
	Letter byte
	Label  string
}

const (
	maxDrives       = 26
	maxPath         = 260
	monitoringDelay = 1 * time.Second

	DriveUnknown   = 0
	DriveNoRootDir = 1
	DriveRemovable = 2
	DriveFixed     = 3
	DriveRemote    = 4
	DriveCDROM     = 5
	DriveRAMDisk   = 6
)

func getLogicalDrives() uint32 {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	ret, _, _ := getLogicalDrives.Call()
	return uint32(ret)
}

func getDriveType(root string) uint32 {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getDriveType := kernel32.MustFindProc("GetDriveTypeA")
	ret, _, _ := getDriveType.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(root))))
	return uint32(ret)
}

func getVolumeLabel(root string) string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getVolumeInformation := kernel32.MustFindProc("GetVolumeInformationA")

	volumeNameBuffer := make([]byte, maxPath+1)

	_, _, _ = getVolumeInformation.Call(
		uintptr(unsafe.Pointer(syscall.StringBytePtr(root))),
		uintptr(unsafe.Pointer(&volumeNameBuffer[0])),
		uintptr(maxPath),
		0,
		0,
		0,
		0,
		0,
	)

	label := string(volumeNameBuffer)
	return strings.Trim(label, "\x00")
}

func findDrive(list []DriveInfo, letter byte) int {
	for i, d := range list {
		if d.Letter == letter {
			return i
		}
	}
	return -1
}

func MonitorUSB(port int) {
	portStr := strconv.Itoa(port)
	var currentDrives []DriveInfo

	for {
		var newList []DriveInfo
		drives := getLogicalDrives()

		for i := 0; i < maxDrives; i++ {
			if (drives & (1 << i)) != 0 {
				letter := byte('A' + i)
				root := fmt.Sprintf("%c:\\", letter)

				if getDriveType(root) == DriveRemovable {
					label := getVolumeLabel(root)
					newList = append(newList, DriveInfo{
						Letter: letter,
						Label:  label,
					})
				}
			}
		}

		for _, d := range newList {
			if findDrive(currentDrives, d.Letter) == -1 {
				fmt.Printf("[+] Connected: %c: — \"%s\"\n", d.Letter, d.Label)
				checkAndRunShell(portStr)
			}
		}

		for _, d := range currentDrives {
			if findDrive(newList, d.Letter) == -1 {
				fmt.Printf("[-] Disabled: %c: — \"%s\"\n", d.Letter, d.Label)
			}
		}

		currentDrives = newList
		time.Sleep(monitoringDelay)
	}
}

func checkAndRunShell(portStr string) {
	exePath, err := os.Executable()
	if err != nil {
		return
	}
	exeDir := filepath.Dir(exePath)

	configPath := filepath.Join(exeDir, "..", "..", "..", "..", "core", "data", "config.json")
	configPath = filepath.Clean(configPath)

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	var config struct {
		Shell int `json:"shell"`
	}

	if err := json.Unmarshal(configFile, &config); err != nil {
		return
	}

	if config.Shell == 0 {
		go runShellWeb(exeDir, portStr)
	}
}

func runShellWeb(baseDir string, portStr string) {
	nm1Path := filepath.Join(baseDir, "..", "..", "..", "..", "shell_NM", "NM1")
	nm1Path = filepath.Clean(nm1Path)

	if err := os.Chdir(nm1Path); err != nil {
		return
	}

	args := []string{
		"dddd",
		"434",
		"555",
		fmt.Sprintf("%s/scan", portStr),
	}

	cmd := exec.Command("./shell_web.exe", args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("error shell_web.exe: %v\n", err)
		return
	}
}

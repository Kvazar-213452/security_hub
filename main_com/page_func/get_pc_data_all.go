package page_func

import (
	"bufio"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")
)

const (
	UNLEN = 256
)

type MEMORYSTATUSEX struct {
	dwLength           uint32
	dwMemoryLoad       uint32
	ullTotalPhys       uint64
	ullAvailPhys       uint64
	ullTotalPageFile   uint64
	ullAvailPageFile   uint64
	ullTotalVirtual    uint64
	ullAvailVirtual    uint64
	ullExtendedVirtual uint64
}

type SYSTEM_INFO struct {
	dwOemId                     uint32
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorArchitecture      uint16
	wReserved                   uint16
}

type OSVERSIONINFO struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]byte
}

// Convert a byte slice to string
func byteSliceToString(b []byte) string {
	return string(b[:len(b)-1]) // Remove the null terminator
}

// GetSystemMemory returns system memory information as a string
func GetSystemMemory() string {
	var memInfo MEMORYSTATUSEX
	memInfo.dwLength = uint32(unsafe.Sizeof(memInfo))

	r1, _, _ := modkernel32.NewProc("GlobalMemoryStatusEx").Call(uintptr(unsafe.Pointer(&memInfo)))
	if r1 == 0 {
		return "Failed to get memory status"
	}

	return fmt.Sprintf("%d MB\n%d MB\n",
		memInfo.ullTotalPhys/(1024*1024),
		memInfo.ullAvailPhys/(1024*1024))
}

// GetProcessorInfo returns processor information as a string
func GetProcessorInfo() string {
	var sysInfo SYSTEM_INFO
	modkernel32.NewProc("GetSystemInfo").Call(uintptr(unsafe.Pointer(&sysInfo)))

	return fmt.Sprintf("%d\n%d\n",
		sysInfo.dwNumberOfProcessors,
		sysInfo.wProcessorArchitecture)
}

// GetOSVersion returns the OS version as a string
func GetOSVersion() string {
	var osvi OSVERSIONINFO
	osvi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	modkernel32.NewProc("GetVersionExW").Call(uintptr(unsafe.Pointer(&osvi)))

	return fmt.Sprintf("%d.%d\n%d\n",
		osvi.dwMajorVersion,
		osvi.dwMinorVersion,
		osvi.dwBuildNumber)
}

// GetComputerNameCustom returns the computer name as a string
func GetComputerNameCustom() string {
	var computerName [syscall.MAX_COMPUTERNAME_LENGTH + 1]byte
	var size uint32 = uint32(len(computerName))

	modkernel32.NewProc("GetComputerNameA").Call(uintptr(unsafe.Pointer(&computerName[0])), uintptr(unsafe.Pointer(&size)))

	return fmt.Sprintf("%s\n", byteSliceToString(computerName[:size]))
}

// GetUserNameCustom returns the username as a string
func GetUserNameCustom() string {
	var userName [UNLEN + 1]byte
	var size uint32 = uint32(len(userName))

	modadvapi32.NewProc("GetUserNameA").Call(uintptr(unsafe.Pointer(&userName[0])), uintptr(unsafe.Pointer(&size)))

	return fmt.Sprintf("%s\n", byteSliceToString(userName[:size]))
}

// GetSystemUptime returns the system uptime as a string
func GetSystemUptime() string {
	getTickCount64 := modkernel32.NewProc("GetTickCount64")

	uptime, _, _ := getTickCount64.Call()
	uptime /= 1000 // Convert from milliseconds to seconds
	days := uptime / (24 * 3600)
	uptime %= 24 * 3600
	hours := uptime / 3600
	uptime %= 3600
	minutes := uptime / 60
	seconds := uptime % 60

	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds\n",
		days, hours, minutes, seconds)
}

func get_usb_info() {
	cmd := exec.Command("wmic", "path", "Win32_PnPEntity", "get", "Name")

	output, err := cmd.Output()
	if err != nil {
		return
	}

	file, err := os.Create(config_main.Devices)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && line != "Name" {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}
}

func Usb_info() string {
	get_usb_info()

	content, err := ioutil.ReadFile(config_main.Devices)
	if err != nil {
		log.Fatalf("Failed to read devices.log: %v", err)
	}

	return string(content)
}

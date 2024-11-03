package page_func

import (
	"fmt"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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

func byteSliceToString(b []byte) string {
	return string(b[:len(b)-1])
}

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

func GetProcessorInfo() string {
	var sysInfo SYSTEM_INFO
	modkernel32.NewProc("GetSystemInfo").Call(uintptr(unsafe.Pointer(&sysInfo)))

	return fmt.Sprintf("%d\n%d\n",
		sysInfo.dwNumberOfProcessors,
		sysInfo.wProcessorArchitecture)
}

func GetOSVersion() string {
	var osvi OSVERSIONINFO
	osvi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	modkernel32.NewProc("GetVersionExW").Call(uintptr(unsafe.Pointer(&osvi)))

	return fmt.Sprintf("%d.%d\n%d\n",
		osvi.dwMajorVersion,
		osvi.dwMinorVersion,
		osvi.dwBuildNumber)
}

func GetComputerNameCustom() string {
	var computerName [syscall.MAX_COMPUTERNAME_LENGTH + 1]byte
	var size uint32 = uint32(len(computerName))

	modkernel32.NewProc("GetComputerNameA").Call(uintptr(unsafe.Pointer(&computerName[0])), uintptr(unsafe.Pointer(&size)))

	return fmt.Sprintf("%s\n", byteSliceToString(computerName[:size]))
}

func GetUserNameCustom() string {
	var userName [UNLEN + 1]byte
	var size uint32 = uint32(len(userName))

	modadvapi32.NewProc("GetUserNameA").Call(uintptr(unsafe.Pointer(&userName[0])), uintptr(unsafe.Pointer(&size)))

	return fmt.Sprintf("%s\n", byteSliceToString(userName[:size]))
}

func GetSystemUptime() string {
	getTickCount64 := modkernel32.NewProc("GetTickCount64")

	uptime, _, _ := getTickCount64.Call()
	uptime /= 1000
	days := uptime / (24 * 3600)
	uptime %= 24 * 3600
	hours := uptime / 3600
	uptime %= 3600
	minutes := uptime / 60
	seconds := uptime % 60

	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds\n",
		days, hours, minutes, seconds)
}

var (
	user32                   = windows.NewLazySystemDLL("user32.dll")
	procEnumWindows          = user32.NewProc("EnumWindows")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")

	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")
)

func isValidProgramWindow(title string) bool {
	title = strings.TrimSpace(title)
	if title == "" {
		return false
	}

	systemWords := []string{
		"Default IME", "MSCTFIME UI", "Program Manager", "Task Host Window",
		"DWM Notification Window", "Service", "Wnd", "Monitor", "Notification",
		"Input", "Webview", "Tray", "Window", "Cmd.exe", "Flyout",
		"Settings", "Indicator", "Host", "UxdService", "NvSvc", "Provider",
	}
	for _, word := range systemWords {
		if strings.Contains(title, word) {
			return false
		}
	}

	systemRegex := regexp.MustCompile(`(?i)(window|helper|service|icon|notification|task|tray|cmd\.exe|\.exe)$`)
	return !systemRegex.MatchString(title)
}

func enumWindows(callback func(hwnd syscall.Handle) bool) {
	cb := syscall.NewCallback(func(hwnd syscall.Handle, lparam uintptr) uintptr {
		if callback(hwnd) {
			return 1
		}
		return 0
	})
	procEnumWindows.Call(cb, 0)
}

func App_open() string {
	content := ""
	windowTitles := make(map[string]bool)

	enumWindows(func(hwnd syscall.Handle) bool {
		length, _, _ := procGetWindowTextLengthW.Call(uintptr(hwnd))
		if length > 0 {
			buffer := make([]uint16, length+1)
			procGetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
			windowTitle := syscall.UTF16ToString(buffer)

			if isValidProgramWindow(windowTitle) {
				if _, exists := windowTitles[windowTitle]; !exists {
					windowTitles[windowTitle] = true
					content += windowTitle + "\n"
				}
			}
		}
		return true
	})

	return string(content)
}

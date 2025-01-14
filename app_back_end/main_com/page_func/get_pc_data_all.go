package page_func

import (
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/shirou/gopsutil/host"
	"github.com/yusufpapurcu/wmi"
	"golang.org/x/sys/windows"
)

var (
	user32                   = windows.NewLazySystemDLL("user32.dll")
	procEnumWindows          = user32.NewProc("EnumWindows")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")

	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	getSystemTimesProc   = kernel32.NewProc("GetSystemTimes")
	psapi                = syscall.NewLazyDLL("psapi.dll")
	getProcessMemoryInfo = psapi.NewProc("GetProcessMemoryInfo")
)

type Win32_BIOS struct {
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	ReleaseDate  string `json:"release_date"`
	SerialNumber string `json:"serial_number"`
}

type Win32_OperatingSystem struct {
	Caption        string `json:"caption"`
	Version        string `json:"version"`
	BuildNumber    string `json:"build_number"`
	OSArchitecture string `json:"os_architecture"`
	SerialNumber   string `json:"serial_number"`
	Manufacturer   string `json:"manufacturer"`
	InstallDate    string `json:"install_date"`
	LastBootUpTime string `json:"last_boot_up_time"`
}

type SystemInfo struct {
	OS              string                  `json:"os"`
	Architecture    string                  `json:"architecture"`
	NumCPU          int                     `json:"num_cpu"`
	HostInfo        *hostInfo               `json:"host_info"`
	BIOSInfo        []Win32_BIOS            `json:"bios_info"`
	OperatingSystem []Win32_OperatingSystem `json:"operating_system_info"`
}

type hostInfo struct {
	Hostname        string `json:"hostname"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	PlatformFamily  string `json:"platform_family"`
	KernelArch      string `json:"kernel_arch"`
	Uptime          uint64 `json:"uptime"`
}

func Get_data_os() (*SystemInfo, error) {
	systemInfo := &SystemInfo{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
	}

	info, _ := host.Info()
	systemInfo.HostInfo = &hostInfo{
		Hostname:        info.Hostname,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		PlatformFamily:  info.PlatformFamily,
		KernelArch:      info.KernelArch,
		Uptime:          info.Uptime,
	}

	var bios []Win32_BIOS
	wmi.Query("SELECT Manufacturer, Name, Version, ReleaseDate, SerialNumber FROM Win32_BIOS", &bios)
	systemInfo.BIOSInfo = bios

	var osInfo []Win32_OperatingSystem
	wmi.Query("SELECT Caption, Version, BuildNumber, OSArchitecture, SerialNumber, Manufacturer, InstallDate, LastBootUpTime FROM Win32_OperatingSystem", &osInfo)
	systemInfo.OperatingSystem = osInfo

	return systemInfo, nil
}

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

type Filetime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

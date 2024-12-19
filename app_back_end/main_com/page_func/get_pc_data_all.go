package page_func

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

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

type SystemInfo struct {
	XMLName        xml.Name `xml:"SystemInfo"`
	Architecture   string   `xml:"Architecture"`
	ProcessorCount int      `xml:"ProcessorCount"`
	OS             struct {
		Name    string `xml:"Name"`
		Version string `xml:"Version"`
	} `xml:"OS"`
	Memory struct {
		FreeMemory        int64 `xml:"FreeMemory"`
		TotalMemory       int64 `xml:"TotalMemory"`
		FreeVirtualMemory int64 `xml:"FreeVirtualMemory"`
	} `xml:"Memory"`
	SystemUptime struct {
		Days    int `xml:"Days"`
		Hours   int `xml:"Hours"`
		Minutes int `xml:"Minutes"`
		Seconds int `xml:"Seconds"`
	} `xml:"SystemUptime"`
	Disk struct {
		FreeSpace  int64 `xml:"FreeSpace"`
		TotalSpace int64 `xml:"TotalSpace"`
	} `xml:"Disk"`
	NetworkAdapters struct {
		Adapters []struct {
			Description string `xml:"Description"`
			IPAddress   string `xml:"IPAddress"`
		} `xml:"Adapter"`
	} `xml:"NetworkAdapters"`
	LoadedLibraries struct {
		Libraries []string `xml:"Library"`
	} `xml:"LoadedLibraries"`
}

func Get_data_os() string {
	exePath := config_main.System_info_exe
	workingDir := config_main.Library_folder
	dataFilePath := "./" + config_main.Library_folder + "/data/" + config_main.File_2_exe_data

	cmd := exec.Command(exePath)
	cmd.Dir = workingDir

	cmd.Run()

	data, _ := ioutil.ReadFile(dataFilePath)

	var sysInfo SystemInfo
	xml.Unmarshal(data, &sysInfo)

	jsonData, _ := json.MarshalIndent(sysInfo, "", "  ")
	func_all.Clear_file(config_main.Global_phat + "\\" + config_main.Library_folder + "\\data\\" + config_main.File_2_exe_data)

	return string(jsonData)
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

//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now
//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now
//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now//data_now

func getCPUUsage() string {
	var idleTime, kernelTime, userTime Filetime

	ret, _, err := getSystemTimesProc.Call(
		uintptr(unsafe.Pointer(&idleTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)
	if ret == 0 {
		return fmt.Sprintf("Failed to get system times: %v", err)
	}

	idle := uint64(idleTime.HighDateTime)<<32 | uint64(idleTime.LowDateTime)
	kernel := uint64(kernelTime.HighDateTime)<<32 | uint64(kernelTime.LowDateTime)
	user := uint64(userTime.HighDateTime)<<32 | uint64(userTime.LowDateTime)

	totalTime := kernel + user
	cpuUsage := (1.0 - float64(idle)/float64(totalTime)) * 100.0

	return fmt.Sprintf("%.2f%%", cpuUsage)
}

func getMemoryUsage() string {
	var memCounters struct {
		cb                         uint32
		PageFaultCount             uint32
		PeakWorkingSetSize         uintptr
		WorkingSetSize             uintptr
		QuotaPeakPagedPoolUsage    uintptr
		QuotaPagedPoolUsage        uintptr
		QuotaPeakNonPagedPoolUsage uintptr
		QuotaNonPagedPoolUsage     uintptr
		PagefileUsage              uintptr
		PeakPagefileUsage          uintptr
	}

	handle, _ := syscall.GetCurrentProcess()

	ret, _, err := getProcessMemoryInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&memCounters)),
		uintptr(unsafe.Sizeof(memCounters)),
	)
	if ret == 0 {
		return fmt.Sprintf("Failed to get memory info: %v", err)
	}

	memoryUsageMB := float64(memCounters.WorkingSetSize) / (1024 * 1024)
	return fmt.Sprintf("%.2f MB", memoryUsageMB)
}

func Get_all_data_now() string {
	cpuInfo := getCPUUsage()
	memoryInfo := getMemoryUsage()

	return fmt.Sprintf("%s\n%s", cpuInfo, memoryInfo)
}

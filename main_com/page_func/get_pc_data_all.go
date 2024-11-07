package page_func

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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
	exePath := "./system_info.exe"
	workingDir := "library"
	dataFilePath := "./library/data/file_2.txt"

	cmd := exec.Command(exePath)
	cmd.Dir = workingDir

	if err := cmd.Run(); err != nil {
		log.Fatalf("Не вдалося запустити exe файл: %v\n", err)
	}

	data, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		log.Fatalf("Не вдалося прочитати файл: %v\n", err)
	}

	var sysInfo SystemInfo
	err = xml.Unmarshal(data, &sysInfo)
	if err != nil {
		log.Fatalf("Не вдалося розпарсити XML: %v", err)
	}

	jsonData, err := json.MarshalIndent(sysInfo, "", "  ")
	if err != nil {
		log.Fatalf("Не вдалося конвертувати в JSON: %v", err)
	}

	return string(jsonData)
}

var (
	user32                   = windows.NewLazySystemDLL("user32.dll")
	procEnumWindows          = user32.NewProc("EnumWindows")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")
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

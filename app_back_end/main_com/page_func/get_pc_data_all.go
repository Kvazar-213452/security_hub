package page_func

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
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
	// Виконання команд для збору інформації про систему
	cmd := exec.Command("uname", "-a")
	osInfo, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running uname: %v", err)
	}

	// Отримання інформації про пам'ять
	cmd = exec.Command("free", "-b")
	memInfo, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running free: %v", err)
	}

	// Отримання часу роботи системи
	cmd = exec.Command("uptime", "-p")
	uptimeInfo, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running uptime: %v", err)
	}

	// Створюємо структуровану інформацію для JSON
	sysInfo := SystemInfo{}
	sysInfo.OS.Name = string(osInfo)
	sysInfo.Memory.FreeMemory = int64(memInfo[0]) // Замість цього, треба парсити результат команди
	sysInfo.SystemUptime.Seconds = int(uptimeInfo[0])

	// Маршалимо в JSON
	jsonData, _ := json.MarshalIndent(sysInfo, "", "  ")
	return string(jsonData)
}

// Функція для визначення, чи є програма у системі
func isValidProgramWindow(title string) bool {
	title = strings.TrimSpace(title)
	if title == "" {
		return false
	}

	systemWords := []string{
		"System", "Task", "Window", "Notification",
	}
	for _, word := range systemWords {
		if strings.Contains(title, word) {
			return false
		}
	}

	systemRegex := regexp.MustCompile(`(?i)(window|helper|service|icon|notification)$`)
	return !systemRegex.MatchString(title)
}

// Збирання інформації про програми
func App_open() string {
	content := ""
	// Тут можна використати команду для перегляду відкритих вікон
	cmd := exec.Command("wmctrl", "-l")
	windowTitles, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running wmctrl: %v", err)
	}

	// Парсимо вивід та перевіряємо
	lines := strings.Split(string(windowTitles), "\n")
	for _, line := range lines {
		if isValidProgramWindow(line) {
			content += line + "\n"
		}
	}

	return content
}

func getCPUUsage() string {
	cmd := exec.Command("top", "-bn1", "|", "grep", "\"Cpu(s)\"")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error getting CPU usage: %v", err)
	}

	// Парсимо інформацію про CPU з виводу команди top
	usage := strings.Split(string(output), ",")[0]
	return fmt.Sprintf("CPU Usage: %s", usage)
}

func getMemoryUsage() string {
	cmd := exec.Command("free", "-m")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error getting memory usage: %v", err)
	}

	// Аналізуємо вивід команди для отримання пам'яті
	memInfo := strings.Split(string(output), "\n")[1]
	memoryUsage := strings.Fields(memInfo)[2]
	return fmt.Sprintf("Memory Usage: %s MB", memoryUsage)
}

func Get_all_data_now() string {
	cpuInfo := getCPUUsage()
	memoryInfo := getMemoryUsage()

	return fmt.Sprintf("%s\n%s", cpuInfo, memoryInfo)
}

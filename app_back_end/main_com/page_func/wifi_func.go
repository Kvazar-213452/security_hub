package page_func

import (
	"encoding/xml"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io/ioutil"
	"os/exec"
	"strings"
)

type WifiInfo struct {
	ProfileName      string   `json:"profile_name"`      // Ім'я профілю
	Version          string   `json:"version"`           // Версія
	Type             string   `json:"type"`              // Тип
	ControlOptions   string   `json:"control_options"`   // Параметри контролю
	SSIDName         string   `json:"ssid_name"`         // Ім'я SSID
	NetworkType      string   `json:"network_type"`      // Тип мережі
	RadioType        string   `json:"radio_type"`        // Тип радіо
	VendorExtension  string   `json:"vendor_extension"`  // Вендорська розширення
	Authentication   string   `json:"authentication"`    // Аутентифікація
	Ciphers          []string `json:"ciphers"`           // Шифри
	SecurityKey      string   `json:"security_key"`      // Ключ безпеки
	KeyContent       string   `json:"key_content"`       // Зміст ключа
	Cost             string   `json:"cost"`              // Вартість
	Congested        string   `json:"congested"`         // Переповненість
	ApproachingLimit string   `json:"approaching_limit"` // Підходження до ліміту даних
	OverLimit        string   `json:"over_limit"`        // Перевищення ліміту даних
	Roaming          string   `json:"roaming"`           // Роумінг
	CostSource       string   `json:"cost_source"`       // Джерело вартості
}

type WifiNetwork struct {
	SSID   string `json:"ssid"`
	Signal int    `json:"signal"`
}

type NetworkInterfaces1 struct {
	Interfaces []NetworkInterface1 `xml:"Interface"`
}

type NetworkInterface1 struct {
	Name            string `xml:"Name"`
	Description     string `xml:"Description"`
	Status          string `xml:"Status"`
	BytesSent       int64  `xml:"BytesSent"`
	BytesReceived   int64  `xml:"BytesReceived"`
	PacketsSent     int64  `xml:"PacketsSent"`
	PacketsReceived int64  `xml:"PacketsReceived"`
}

type WifiNetwork_1 struct {
	SSID          string `xml:"SSID"`
	SignalQuality int    `xml:"SignalQuality"`
}

type Networks struct {
	XMLName  xml.Name        `xml:"Networks"`
	Networks []WifiNetwork_1 `xml:"Network"`
}

func Get_Wifi_info() (*WifiInfo, error) {
	// Використовуємо nmcli для отримання підключених мереж
	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SECURITY", "device", "wifi", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("не вдалося отримати інформацію про мережу: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	wifiInfo := &WifiInfo{}

	// Проходимо по кожному рядку та перевіряємо наявність потрібних даних
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				if strings.Contains(parts[0], "SSID") {
					wifiInfo.SSIDName = strings.TrimSpace(parts[1])
				} else if strings.Contains(parts[0], "SECURITY") {
					// Тут можна обробляти безпеку
					wifiInfo.Authentication = strings.TrimSpace(parts[1])
				}
			}
		}
	}

	// Перевірка на наявність SSIDName
	if wifiInfo.SSIDName == "" {
		return nil, fmt.Errorf("інформацію про Wi-Fi не знайдено")
	}

	return wifiInfo, nil
}

func Get_available_Wifi_networks() ([]WifiNetwork, error) {
	// Run iwlist to scan networks on wlan0 (or your network interface)
	cmd := exec.Command("sudo", "iwlist", "wlan0", "scan")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running iwlist command: %v", err)
	}

	// Parse the output of the iwlist scan
	var networks []WifiNetwork
	var ssid string
	var signalQuality int

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "ESSID:") {
			// Extract SSID
			ssid = strings.Trim(strings.Split(line, ":")[1], "\"")
		}
		if strings.Contains(line, "Signal level=") {
			// Extract signal strength
			_, err := fmt.Sscanf(line, " Signal level=%d", &signalQuality)
			if err == nil {
				// Add the Wi-Fi network to the list
				networks = append(networks, WifiNetwork{
					SSID:   ssid,
					Signal: signalQuality,
				})
			}
		}
	}

	// Return the JSON data
	return networks, nil
}

type Network struct {
	XMLName xml.Name `xml:"Networks"`
	SSIDs   []string `xml:"SSID"`
}

func Get_connected_SSID() string {
	exePath := config_main.Get_ssid_exe
	workingDir := config_main.Library_folder
	dataFilePath := "./" + config_main.Library_folder + "/data/" + config_main.File_exe_data

	cmd := exec.Command(exePath)
	cmd.Dir = workingDir

	cmd.Run()
	data, _ := ioutil.ReadFile(dataFilePath)

	var network Network
	xml.Unmarshal(data, &network)

	func_all.Clear_file(config_main.Global_phat + "\\" + config_main.Library_folder + "\\data\\" + config_main.File_exe_data)

	if len(network.SSIDs) > 0 {
		return network.SSIDs[0]
	}
	return ""
}

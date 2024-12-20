package page_func

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
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
	Signal string `json:"signal"`
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
	ssid := Get_connected_SSID()

	cmd := exec.Command("netsh", "wlan", "show", "profile", "name="+ssid, "key=clear")
	output, _ := cmd.CombinedOutput()

	lines := strings.Split(string(output), "\n")
	wifiInfo := &WifiInfo{}

	for _, line := range lines {
		if strings.Contains(line, "Profile") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.ProfileName = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Version") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Version = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Type") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Type = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Control options") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.ControlOptions = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "SSID name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.SSIDName = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Network type") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.NetworkType = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Radio type") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.RadioType = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Vendor extension") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.VendorExtension = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Authentication") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Authentication = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Cipher") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Ciphers = append(wifiInfo.Ciphers, strings.TrimSpace(parts[1]))
			}
		} else if strings.Contains(line, "Security key") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.SecurityKey = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Key Content") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.KeyContent = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Cost") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Cost = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Congested") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Congested = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Approaching Data Limit") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.ApproachingLimit = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Over Data Limit") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.OverLimit = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Roaming") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Roaming = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Cost Source") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.CostSource = strings.TrimSpace(parts[1])
			}
		}
	}

	if wifiInfo.SSIDName == "" {
		return nil, fmt.Errorf("інформацію про Wi-Fi не знайдено")
	}

	return wifiInfo, nil
}

func Get_available_Wifi_networks() ([]WifiNetwork_1, error) {
	exePath := config_main.Available_wifi_exe
	workingDir := config_main.Library_folder
	dataFilePath := "./" + config_main.Library_folder + "/data/" + config_main.File_1_exe_data

	cmd := exec.Command(exePath)
	cmd.Dir = workingDir
	cmd.Run()

	data, _ := ioutil.ReadFile(dataFilePath)

	var networks Networks
	xml.Unmarshal(data, &networks)

	func_all.Clear_file(config_main.Global_phat + "\\" + config_main.Library_folder + "\\data\\" + config_main.File_1_exe_data)

	return networks.Networks, nil
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

func Get_info_packages_wifi() []byte {
	cmd := exec.Command(config_main.Wifi_packege_data_exe)
	cmd.Dir = config_main.Library_folder
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Run()

	filePath := "./" + config_main.Library_folder + config_main.File_data_exe_wifi_packege
	xmlFile, _ := os.Open(filePath)

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var networkInterfaces NetworkInterfaces1

	xml.Unmarshal(byteValue, &networkInterfaces)
	jsonData, _ := json.MarshalIndent(networkInterfaces, "", "  ")

	return jsonData
}

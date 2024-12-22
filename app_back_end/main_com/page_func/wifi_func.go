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
	"syscall"
)

type WiFiData struct {
	XMLName        xml.Name `xml:"WiFiData"`
	Name           string   `xml:"Name"`
	Description    string   `xml:"Description"`
	GUID           string   `xml:"GUID"`
	State          string   `xml:"State"`
	SignalStrength string   `xml:"SignalStrength"`
	RadioType      string   `xml:"RadioType"`
	BSSID          string   `xml:"BSSID"`
	Frequency      string   `xml:"Frequency"`
	Channel        string   `xml:"Channel"`
	SSID           string   `xml:"SSID"`
	Authentication string   `xml:"Authentication"`
	Cipher         string   `xml:"Cipher"`
	ConnectionMode string   `xml:"ConnectionMode"`
	ProfileType    string   `xml:"ProfileType"`
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

func Get_Wifi_info() (*WiFiData, error) {
	exePath := "./MyConsoleApp.exe"
	workingDir := "./library/get_netsh"
	dataFilePath := "./library/get_netsh/" + "main.xml"

	cmd := exec.Command(exePath)
	cmd.Dir = workingDir
	cmd.Run()

	data, _ := ioutil.ReadFile(dataFilePath)

	var wifiData WiFiData
	xml.Unmarshal(data, &wifiData)

	if wifiData.SSID == "" {
		return nil, fmt.Errorf("інформацію про Wi-Fi не знайдено")
	}

	return &wifiData, nil
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

	filePath := "./" + config_main.Library_folder + "/" + config_main.File_data_exe_wifi_packege
	xmlFile, _ := os.Open(filePath)

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var networkInterfaces NetworkInterfaces1

	xml.Unmarshal(byteValue, &networkInterfaces)
	jsonData, _ := json.MarshalIndent(networkInterfaces, "", "  ")

	return jsonData
}

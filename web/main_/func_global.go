package main_

import (
	"fmt"
	"os/exec"
	"strings"
)

type WifiInfo struct {
	SSID            string `json:"ssid"`
	Description     string `json:"description"`
	GUID            string `json:"guid"`
	PhysicalAddress string `json:"physical_address"`
	State           string `json:"state"`
	RadioType       string `json:"radio_type"`
	Authentication  string `json:"authentication"`
	SignalStrength  string `json:"signal_strength"`
}

type WifiNetwork struct {
	SSID   string `json:"ssid"`
	Signal string `json:"signal"`
}

func Get_Wifi_info() (*WifiInfo, error) {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("не вдалося виконати команду: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	wifiInfo := &WifiInfo{}

	for _, line := range lines {
		if strings.Contains(line, "SSID") && !strings.Contains(line, "BSSID") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.SSID = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Description") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Description = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "GUID") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.GUID = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Physical address") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.PhysicalAddress = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "State") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.State = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Radio type") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.RadioType = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Authentication") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				wifiInfo.Authentication = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "Signal") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				signal := strings.TrimSpace(parts[1])
				wifiInfo.SignalStrength = strings.Replace(signal, "%", "", -1)
			}
		}
	}

	if wifiInfo.SSID == "" {
		return nil, fmt.Errorf("інформацію про Wi-Fi не знайдено")
	}

	return wifiInfo, nil
}

func Get_available_Wifi_networks() ([]WifiNetwork, error) {
	cmd := exec.Command("netsh", "wlan", "show", "networks", "mode=bssid")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("не вдалося виконати команду: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var networks []WifiNetwork
	var currentNetwork WifiNetwork
	var bssidCount int

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "SSID") {
			if currentNetwork.SSID != "" {
				networks = append(networks, currentNetwork)
				currentNetwork = WifiNetwork{}
				bssidCount = 0
			}
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				currentNetwork.SSID = strings.TrimSpace(parts[1])
			} else {
				currentNetwork.SSID = ""
			}
		}

		if strings.HasPrefix(line, "BSSID") {
			bssidCount++
		}

		if strings.HasPrefix(line, "Signal") {
			if bssidCount == 1 {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					currentNetwork.Signal = strings.TrimSpace(parts[1])
				}
			} else if bssidCount == 2 {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					currentNetwork.Signal = strings.TrimSpace(parts[1])
				}
			}
		}
	}

	if currentNetwork.SSID != "" {
		networks = append(networks, currentNetwork)
	}

	if len(networks) == 0 {
		return nil, fmt.Errorf("доступні мережі не знайдені")
	}

	return networks, nil
}

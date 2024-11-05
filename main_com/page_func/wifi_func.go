package page_func

import (
	"fmt"
	"io/ioutil"
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

func Get_Wifi_info() (*WifiInfo, error) {
	ssid := Get_connected_SSID()

	cmd := exec.Command("netsh", "wlan", "show", "profile", fmt.Sprintf("name=%s", ssid), "key=clear")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("не вдалося виконати команду: %v", err)
	}

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

func Get_connected_SSID() string {
	dll := syscall.MustLoadDLL("library/get_ssid.dll")
	proc := dll.MustFindProc("GetConnectedSSIDAndWriteToFile")

	proc.Call()

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Не вдалося прочитати файл:", err)
		return ""
	}

	ssid := string(data)
	ssid = strings.TrimPrefix(ssid, "Connected SSID: ")
	ssid = strings.TrimSpace(ssid)

	return ssid
}

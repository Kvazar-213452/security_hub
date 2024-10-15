package main_

import (
	"encoding/json"
	"head/main_/func_all"
	"net/http"
	"strconv"
	"strings"
	"syscall"
)

type VisualizationMessage struct {
	Message int `json:"message"`
}

type OSData struct {
	SystemMemory  string `json:"system_memory"`
	ProcessorInfo string `json:"processor_info"`
	OSVersion     string `json:"os_version"`
	ComputerName  string `json:"computer_name"`
	UserName      string `json:"user_name"`
	SystemUptime  string `json:"system_uptime"`
}

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_wifi_now post")

		wifiInfo, err := func_all.Get_Wifi_info()
		if err != nil {
			http.Error(w, "Помилка отримання інформації про Wi-Fi", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wifiInfo)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_wifi_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_wifi post")
		networks, err := func_all.Get_available_Wifi_networks()
		if err != nil {
			http.Error(w, "Помилка отримання інформації про Wi-Fi мережі", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(networks)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_server_fet_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_logs post")

		jsonData, err := func_all.LoadLogFile()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_network_now(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/network_now post")

		ssid := func_all.GetConnectedSSID()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"ssid": ssid}); err != nil {
			http.Error(w, "Помилка при кодуванні JSON", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_config_global(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/config_global post")

		config, err := func_all.LoadConfig()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(config)
		if err != nil {
			http.Error(w, "не вдалося закодувати в JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_config_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/visualization change post")

		var msg VisualizationMessage

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Не вдалося декодувати JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		func_all.UpdateVisualization(strconv.Itoa(msg.Message), "Visualization")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_get_os_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("get os data")

		dll := syscall.NewLazyDLL("library/system_info.dll")

		getSystemMemory := dll.NewProc("GetSystemMemory")
		getProcessorInfo := dll.NewProc("GetProcessorInfo")
		getOSVersion := dll.NewProc("GetOSVersion")
		getComputerNameCustom := dll.NewProc("GetComputerNameCustom")
		getUserNameCustom := dll.NewProc("GetUserNameCustom")
		getSystemUptime := dll.NewProc("GetSystemUptime")

		systemMemory := func_all.RemoveNewlines(func_all.CallDLLFunction(getSystemMemory, "System Memory Info"))
		processorInfo := func_all.RemoveNewlines(func_all.CallDLLFunction(getProcessorInfo, "Processor Info"))
		osVersion := func_all.RemoveNewlines(func_all.CallDLLFunction(getOSVersion, "OS Version Info"))
		computerName := func_all.RemoveNewlines(func_all.CallDLLFunction(getComputerNameCustom, "Computer Name Info"))
		userName := func_all.RemoveNewlines(func_all.CallDLLFunction(getUserNameCustom, "User Name Info"))
		systemUptime := func_all.RemoveNewlines(func_all.CallDLLFunction(getSystemUptime, "System Uptime Info"))

		osData := OSData{
			SystemMemory:  systemMemory,
			ProcessorInfo: processorInfo,
			OSVersion:     osVersion,
			ComputerName:  computerName,
			UserName:      userName,
			SystemUptime:  systemUptime,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(osData)
		if err != nil {
			http.Error(w, "Помилка при формуванні JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_usb_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("usb info")

		info := func_all.Usb_info()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"devices": devices}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Не вдалося кодувати JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_resource_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("resource info")

		info := func_all.Resource_info()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"devices": devices}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Не вдалося кодувати JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

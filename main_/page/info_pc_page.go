package page

import (
	"encoding/json"
	"head/main_/func_all"
	"head/main_/page_func"
	"net/http"
	"strings"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type OSData struct {
	SystemMemory  string `json:"system_memory"`
	ProcessorInfo string `json:"processor_info"`
	OSVersion     string `json:"os_version"`
	ComputerName  string `json:"computer_name"`
	UserName      string `json:"user_name"`
	SystemUptime  string `json:"system_uptime"`
}

func Post_get_os_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("get os data")

		systemMemory := page_func.GetSystemMemory()
		processorInfo := page_func.GetProcessorInfo()
		osVersion := page_func.GetOSVersion()
		computerName := page_func.GetComputerNameCustom()
		userName := page_func.GetUserNameCustom()
		systemUptime := page_func.GetSystemUptime()

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

		info := page_func.Usb_info()
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

		info := page_func.Get_all_data_now()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"data": devices}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Не вдалося кодувати JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

package page

import (
	"encoding/json"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"net/http"
	"strings"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type OSData struct {
	Data string `json:"data"`
}

func Post_get_os_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("get os data")

		data := page_func.Get_data_os()

		osData := OSData{
			Data: data,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(osData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_window_open(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("usb info")

		info := page_func.App_open()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"devices": devices}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
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
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

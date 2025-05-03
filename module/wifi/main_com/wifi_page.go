package main_com

import (
	"encoding/json"
	"fmt"
	"head/main_com/speed_count"
	"head/main_com/wifi"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// module/wifi/main_com/wifi_page.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		wifiInfo, err := wifi.Get_Wifi_info()
		if err != nil {
			http.Error(w, "Помилка отримання інформації про Wi-Fi", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wifiInfo)
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

func Post_wifi_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		networks, err := wifi.Get_available_Wifi_networks()
		if err != nil {
			json.NewEncoder(w).Encode("error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(networks)
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

func Post_network_now(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		ssid := wifi.Get_connected_SSID()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"ssid": ssid}); err != nil {
			http.Error(w, "Помилка при кодуванні JSON", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

func Post_get_pacage_info_wifi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		jsonData := wifi.Get_info_packages_wifi()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(jsonData))
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

func Get_file(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	var requestData map[string]interface{}
	json.Unmarshal(body, &requestData)

	filePath := filepath.Join(requestData["data"].(string))

	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func Get_speed(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	downloadSpeed, uploadSpeed := speed_count.Start_speed_count()

	speeds := []string{
		fmt.Sprintf("%.2f", downloadSpeed),
		fmt.Sprintf("%.2f", uploadSpeed),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(speeds)
}

package main_

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		wifiInfo, err := Get_Wifi_info()
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
		networks, err := Get_available_Wifi_networks()
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
		config_main := LoadConfig_main("config.toml")
		port_main := fmt.Sprintf(":%d", config_main.Port)

		Other_server_post_get_log(w, port_main)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_network_now(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ssid := GetConnectedSSID()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"ssid": ssid}); err != nil {
			http.Error(w, "Помилка при кодуванні JSON", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

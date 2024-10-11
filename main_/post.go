package main_

import (
	"encoding/json"
	"head/main_/func_all"
	"net/http"
	"strconv"
)

type VisualizationMessage struct {
	Message int `json:"message"`
}

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

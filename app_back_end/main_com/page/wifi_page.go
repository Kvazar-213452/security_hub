package page

import (
	"encoding/json"
	"fmt"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"net/http"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_wifi_now post")

		wifiInfo, err := page_func.Get_Wifi_info()
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
		func_all.AppendToLog("/get_wifi post")
		networks, err := page_func.Get_available_Wifi_networks()
		if err != nil {
			json.NewEncoder(w).Encode("error")
			return
		}

		fmt.Println(networks)
		fmt.Println("networks")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(networks)
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

func Post_network_now(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/network_now post")

		ssid := page_func.Get_connected_SSID()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"ssid": ssid}); err != nil {
			http.Error(w, "Помилка при кодуванні JSON", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "error metod", http.StatusMethodNotAllowed)
	}
}

package main_

import (
	"encoding/json"
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

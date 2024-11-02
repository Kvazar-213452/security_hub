package main_com

import (
	"head/main_com/func_all"
	"net/http"
)

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

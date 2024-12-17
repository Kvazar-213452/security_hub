package page

import (
	"encoding/json"
	"fmt"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"net/http"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type folder_info struct {
	Rootsize float64
	Top      [][]string
}

func Post_scan_dir(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/scan_dir post")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Неможливо прочитати тіло запиту", http.StatusInternalServerError)
			return
		}
		fmt.Println("Отримані дані від клієнта:", string(body))

		var request struct {
			Dir  string   `json:"dir"`
			Mas1 []string `json:"mas1"`
			Mas2 []string `json:"mas2"`
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "Неправильний формат запиту", http.StatusBadRequest)
			return
		}

		rootSize, unix := page_func.Run_scan_dir(request.Dir, request.Mas1)

		Folder_info := folder_info{
			Rootsize: rootSize,
			Top:      unix,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Folder_info)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

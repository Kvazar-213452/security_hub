package main_com

import (
	"encoding/json"
	"fmt"
	"head/main_com/system"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// module/system/main_com/system.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type OSData struct {
	Data string `json:"data"`
}

func Post_get_os_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data, _ := system.Get_data_os()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_window_open(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		info := system.App_open()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"devices": devices}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
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

// scan_dir // scan_dir // scan_dir // scan_dir // scan_dir // scan_dir
// scan_dir // scan_dir // scan_dir // scan_dir // scan_dir // scan_dir
// scan_dir // scan_dir // scan_dir // scan_dir // scan_dir // scan_dir

type folder_info struct {
	Rootsize float64
	Top      [][]string
}

func Post_scan_dir(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)

		var request struct {
			Dir  string   `json:"dir"`
			Mas1 []string `json:"mas1"`
			Mas2 []string `json:"mas2"`
		}
		json.Unmarshal(body, &request)

		rootSize, unix := system.Run_scan_dir(request.Dir, request.Mas1)

		Folder_info := folder_info{
			Rootsize: rootSize,
			Top:      unix,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Folder_info)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

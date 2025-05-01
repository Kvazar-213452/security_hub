package main_com

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// app_back_end/main_com/page/cleanup.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type CleanupData struct {
	Backup  int `json:"backup"`
	Wifi    int `json:"wifi"`
	Desktop int `json:"desktop"`
	Doskey  int `json:"doskey"`
}

func Post_cleanup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var dataCleanup CleanupData

		json.NewDecoder(r.Body).Decode(&dataCleanup)

		if dataCleanup.Wifi == 1 {
			Cleanup_wifi()
		} else if dataCleanup.Backup == 1 {
			Cleanup_backup()
		} else if dataCleanup.Desktop == 1 {
			Cleanup_desktop()
		} else if dataCleanup.Doskey == 1 {
			Cleanup_doskey()
		}

		Cleanup()

		w.Write(nil)
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

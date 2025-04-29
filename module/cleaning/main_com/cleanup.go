package main_com

import (
	"encoding/json"
	"head/main_com/func_all"
	"net/http"
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

		func_all.AppendToLog("cleanup")
		Cleanup()

		w.Write(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

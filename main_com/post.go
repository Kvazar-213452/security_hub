package main_com

import (
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io/ioutil"
	"net/http"

	"github.com/pkg/browser"
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

func Post_Browser_site_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_logs post")

		func_all.AppendToLog("transition to /Browser_site_app")

		url := config_main.Site_main
		browser.OpenURL(url)

		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_get_style(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_style post")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data, _ := ioutil.ReadFile("data/style/main.css")

		decodedString := string(data)

		json.NewEncoder(w).Encode(decodedString)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

package main_com

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

func Post_install_style(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/install_style post")

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Не вдалося отримати файл: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		savePath := filepath.Join("data", "style", "main.css")

		outFile, err := os.Create(savePath)
		if err != nil {
			http.Error(w, "Не вдалося створити файл: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "Не вдалося зберегти файл: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_version_get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/version_get post")

		version := func_all.Version_server()

		config, err := func_all.LoadConfig_start(config_main.Main_config)
		if err != nil {
			fmt.Printf("Не вдалося завантажити конфігурацію: %v\n", err)
			return
		}

		type Data_ump struct {
			Version        int
			Version_config int
		}

		Data := Data_ump{
			Version:        version,
			Version_config: config.Version,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Data)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

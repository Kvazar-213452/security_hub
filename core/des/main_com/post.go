package main_com

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/browser"
)

type Data_ump struct {
	Version_config int
}

func Post_version_get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		config := func_all.LoadConfig_start(config_main.Main_config)

		Data := Data_ump{
			Version_config: config.Version,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Data)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_info_work_server_register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		status := func_all.Check_server_Status(config_main.Server_register_and_data_url + config_main.Server_register_and_data_url_check)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": status})
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_info_work_server_data_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		status := func_all.Check_server_Status(config_main.Server_data_file_url + config_main.Server_data_file_url_server_unix)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": status})
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

func Post_Browser_site_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		url := config_main.Site_main
		browser.OpenURL(url)

		w.Write(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_version_get_server(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("ERROR: Invalid request method:", r.Method)
		http.Error(w, "error", http.StatusMethodNotAllowed)
		return
	}

	version := func_all.Get_server_version()

	Data := Data_ump{
		Version_config: version,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Data); err != nil {
		log.Println("ERROR: Failed to encode response:")
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func Post_log_out(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.Clear_file(config_main.Data_user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_install_style(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, _ := r.FormFile("file")
		defer file.Close()

		savePath := filepath.Join("../data", "main.css")

		outFile, _ := os.Create(savePath)
		defer outFile.Close()

		io.Copy(outFile, file)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Rost_open_url(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var request struct {
			Data string `json:"Data"`
		}

		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		url := request.Data
		browser.OpenURL(url)

		w.Write(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

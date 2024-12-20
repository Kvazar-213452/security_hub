package main_com

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/browser"
)

// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠻⢷⣦⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣦⣄⣠⡤⠒⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⢾⣤⣤⣤⣴⣶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠙⠻⣿⣿⣿⣿⣿⣯⣿⣿⣿⣭⣭⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀
// ⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣿⢿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀
// ⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃⣿⣿⣿⣿⣿⣿⣿⠹⣿⣿⣿⣿⣿⣿⠟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠻⠆⠀⠀
// ⠀⠰⠿⠿⠿⣿⣿⣿⣿⣿⣿⡿⠀⢠⣿⣿⣿⣿⣿⣿⣿⡄⣿⣿⣿⣿⣿⣿⠀⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀
// ⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⠑⠊⠙⣿⣿⣿⠿⠿⠿⠛⠃⠙⠛⣛⣿⣿⣿⣄⡸⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀
// ⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⡧⣴⡾⠿⠛⠛⠉⠀⠀⠀⠀⠀⠀⠀⠙⠛⠛⠛⠿⢿⣻⠶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀
// ⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⡅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀
// ⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⡿⠋⠉⠉⠻⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀
// ⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣷⢾⣍⠙⠆⢸⣿⣿⣿⣿⣿⣿⣿⡆⠀
// ⠐⠀⠴⠀⠀⠀⢹⣿⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⠈⠛⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⡧⠴⠋⠀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣷⠀
// ⠀⠀⠦⠄⠀⠀⠀⠹⣿⣿⠿⣧⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
// ⠀⠀⠀⠀⠀⣀⣀⠀⢀⣿⣿⣿⣿⣿⣶⣦⣤⣀⣀⡀⠀⠀⠀⠀⠀⠀⢀⣤⣴⣾⣿⣿⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄
// ⠀⠀⠀⠀⠀⢿⣽⠧⠾⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⡄⠀⠀⠀⢸⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
// ⠀⠀⠀⠀⠀⠀⠁⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⢻⣿⠿⠛⠉⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣾⣿⣿⣿⣿⣿⠃
// ⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⡿⢿⣿⣿⠟⠁⠀⣀⣤⠴⠛⠉⠀⠀⠀⠀⠀⠈⢻⣿⣿⣿⣿⣿⣿⣿⣻⣿⣿⣿⣿⣿⣿⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⠟⠋⣀⣾⣋⣤⡤⠖⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⡇⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⡿⠛⠁⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⣿⣿⡏⢸⣿⡿⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⡉⠉⠉⠉⠀⠈⠉⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢳⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣷⠀⠀⠀⠀⠀⠀⠀⠀

type Data_ump struct {
	Version_config int
}

func Post_server_fet_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_logs post")

		jsonData := func_all.LoadLogFile()

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
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
		http.Error(w, "error", http.StatusMethodNotAllowed)
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
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_install_style(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/install_style post")

		file, _, _ := r.FormFile("file")
		defer file.Close()

		savePath := filepath.Join("data", "style", "main.css")

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

func Post_version_get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/version_get post")

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

func Post_version_get_server(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_version_get_server post")

		req, _ := http.NewRequest("POST", config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_version, bytes.NewBuffer([]byte{}))

		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var response map[string]interface{}

		json.Unmarshal(body, &response)

		versionStr, ok := response["version"].(string)
		if !ok {
			log.Fatal("Version is not a string")
		}

		version, _ := strconv.Atoi(versionStr)

		Data := Data_ump{
			Version_config: version,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Data)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_reg_status(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_reg_status post")

		config := page_func.Get_config_user()

		Data := page_func.Config_reg{
			Name:   config.Name,
			Pasw:   config.Pasw,
			Gmail:  config.Gmail,
			Acsses: config.Acsses,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Data)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_info_work_server_data_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_get_info_work_server_data_file post")

		status := func_all.Check_server_Status(config_main.Server_data_file_url + config_main.Server_data_file_url_server_unix)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": status})
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_info_work_server_register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_get_info_work_server_register post")
		status := func_all.Check_server_Status(config_main.Server_register_and_data_url + config_main.Server_register_and_data_url_check)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": status})
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_log_out(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_log_out post")

		func_all.Clear_file(config_main.Data_user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

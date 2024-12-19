package main_com

import (
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

		type Data_ump struct {
			Version_config int
		}

		Data := Data_ump{
			Version_config: config.Version,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Data)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

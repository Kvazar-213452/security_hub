package main_com

import (
	"encoding/json"
	"head/main_com/func_all"
	"head/main_com/model"
	"head/main_com/settings"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// app_back_end/main_com/page/model_page.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type VisualizationMessage struct {
	Message int `json:"message"`
}

var data struct {
	Value string `json:"value"`
}

type ModuleData struct {
	Module string `json:"module"`
}

func Post_config_global(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		config := LoadConfig()

		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(config)

		w.Write(jsonData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_config_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var msg VisualizationMessage

		json.NewDecoder(r.Body).Decode(&msg)

		UpdateVisualization(strconv.Itoa(msg.Message), "Visualization")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_log_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		UpdateConfigKey("log", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_port_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		UpdateConfigKey("port", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_shell_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		UpdateConfigKey("shell", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_change_lang_model(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		UpdateConfigKey("lang", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_style_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		if data.Value == "1" {
			UpdateConfigKey("style", "main")
		} else {
			UpdateConfigKey("style", "null")
		}

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_updata_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.Updata_app()

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_accses_updata(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		config := func_all.LoadConfig_start("../data/config.json")
		version := func_all.Get_server_version()

		if version > config.Version {
			w.Write([]byte("1"))
		} else {
			w.Write([]byte("0"))
		}
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_info_module_nm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, err := os.Open("../data/module_config.json")
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		var data interface{}
		err = json.Unmarshal(fileData, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_install_module(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data ModuleData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		val := model.Install_NM(data.Module)

		if val == 0 {
			w.Write([]byte("0"))
		} else {
			w.Write([]byte("1"))
		}
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_uninstall_module(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data ModuleData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		val := model.Uninstall_NM(data.Module)

		if val == 0 {
			w.Write([]byte("0"))
		} else {
			w.Write([]byte("1"))
		}
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_del_temp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.Del_temp()

		w.Write([]byte("0"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_change_lang_settings(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		json.NewDecoder(r.Body).Decode(&data)

		UpdateConfigKey("lang", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_temp_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		val := settings.Get_info()

		w.Write([]byte(val))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

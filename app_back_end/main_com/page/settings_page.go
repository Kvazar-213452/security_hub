package page

import (
	"encoding/json"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"net/http"
	"strconv"
)

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

func Post_config_global(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/config_global post")

		config := page_func.LoadConfig()

		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(config)

		w.Write(jsonData)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_config_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/visualization change post")

		var msg VisualizationMessage

		json.NewDecoder(r.Body).Decode(&msg)

		page_func.UpdateVisualization(strconv.Itoa(msg.Message), "Visualization")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_log_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/log_change change post")

		json.NewDecoder(r.Body).Decode(&data)

		page_func.UpdateConfigKey("log", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_port_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/port_change change post")

		json.NewDecoder(r.Body).Decode(&data)

		page_func.UpdateConfigKey("port", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_shell_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/shell_change change post")

		json.NewDecoder(r.Body).Decode(&data)

		page_func.UpdateConfigKey("shell", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_change_lang_settings(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/change_lang_settings post")

		json.NewDecoder(r.Body).Decode(&data)

		page_func.UpdateConfigKey("lang", data.Value)

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_style_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/style_change post")

		json.NewDecoder(r.Body).Decode(&data)

		if data.Value == "1" {
			page_func.UpdateConfigKey("style", "main")
		} else {
			page_func.UpdateConfigKey("style", "null")
		}

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

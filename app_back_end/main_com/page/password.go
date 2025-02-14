package page

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"log"
	"net/http"
)

// app_back_end/main_com/page/password.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type RequestData_o1 struct {
	Gmail string `json:"gmail"`
}

type RequestData_dqwd struct {
	Key  string `json:"key"`
	Pasw string `json:"pasw"`
}

type RequestData_dqwd1 struct {
	Key   string `json:"key"`
	Pasw  string `json:"pasw"`
	Gmail string `json:"gmail"`
}

type RequestData5 struct {
	Data string `json:"data"`
}

func Post_get_password(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_password post")

		config := page_func.Get_config_user()

		data := RequestData_o1{
			Gmail: config.Gmail,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_get_password, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(body))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_add_key_pasw(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/add_key_pasw post")

		var data RequestData_dqwd
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		config := page_func.Get_config_user()

		data1 := RequestData_dqwd1{
			Key:   data.Key,
			Pasw:  data.Pasw,
			Gmail: config.Gmail,
		}

		jsonData, err := json.Marshal(data1)
		if err != nil {
			log.Println("Error JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_add_key_pasw, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(body))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_del_key_pasw(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/del_key_pasw post")

		var requestData RequestData5
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		config := page_func.Get_config_user()

		data1 := RequestData_dqwd{
			Key:  config.Gmail,
			Pasw: requestData.Data,
		}

		jsonData, err := json.Marshal(data1)
		if err != nil {
			log.Println("Error JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_del_key_pasw, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(body))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

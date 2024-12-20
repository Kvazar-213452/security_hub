package page

import (
	"bytes"
	"encoding/json"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io/ioutil"
	"net/http"
)

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

func Post_get_password(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_password post")

		config := page_func.Get_config_user()

		data := RequestData_o1{
			Gmail: config.Gmail,
		}

		jsonData, _ := json.Marshal(data)

		resp, _ := http.Post("http://127.0.0.1:5000/get_password", "application/json", bytes.NewBuffer(jsonData))
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(body))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_add_key_pasw(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/add_key_pasw post")

		var data RequestData_dqwd
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Помилка при обробці JSON", http.StatusBadRequest)
			return
		}

		config := page_func.Get_config_user()

		data1 := RequestData_dqwd1{
			Key:   data.Key,
			Pasw:  data.Pasw,
			Gmail: config.Gmail,
		}

		jsonData, _ := json.Marshal(data1)

		resp, _ := http.Post("http://127.0.0.1:5000/add_key_pasw", "application/json", bytes.NewBuffer(jsonData))
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(body))
	}
}

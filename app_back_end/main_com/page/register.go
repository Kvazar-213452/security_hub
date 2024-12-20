package page

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"io/ioutil"
	"net/http"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type User struct {
	Name  string `json:"name"`
	Pasw  string `json:"pasw"`
	Gmail string `json:"gmail"`
	Code  string `json:"code"`
}

func Post_send_email(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_send_email post")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var request struct {
			Name     string `json:"name"`
			Gmail    string `json:"gmail"`
			Password string `json:"password"`
		}

		json.Unmarshal(body, &request)

		code_ril := page_func.GenerateRandomDigits()

		text_xxx := page_func.Cripter_xxx(request.Gmail)
		сode := page_func.Cripter_xxx(code_ril)

		data := page_func.RequestData_xxx{
			Receiver: text_xxx,
			Code:     сode,
		}

		url := config_main.Server_register_and_data_url + config_main.Server_register_and_data_url_send_email
		page_func.SendPostRequest_xxx(url, data)

		config := page_func.Config_reg{
			Name:   request.Name,
			Pasw:   request.Password,
			Gmail:  request.Gmail,
			Code:   page_func.Encrypt_code_reg_save(code_ril),
			Acsses: "0",
		}

		page_func.Save_data_reg(config)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_code_verefic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_code_verefic post")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var request struct {
			Code string `json:"code"`
		}

		json.Unmarshal(body, &request)

		filePath := config_main.Data_user
		fileContent, _ := ioutil.ReadFile(filePath)

		var user User

		json.Unmarshal(fileContent, &user)

		code_ril := page_func.Decrypt_code_reg_save(user.Code)

		if request.Code == code_ril {
			config := page_func.Get_config_user()

			config1 := page_func.Config_reg{
				Name:   config.Name,
				Pasw:   config.Pasw,
				Gmail:  config.Gmail,
				Code:   config.Code,
				Acsses: config.Acsses,
			}

			status := page_func.Send_user_data_server(config1)

			if status == "1" {
				config2 := page_func.Config_reg{
					Name:   config.Name,
					Pasw:   config.Pasw,
					Gmail:  config.Gmail,
					Code:   config.Code,
					Acsses: "1",
				}

				page_func.Save_data_reg(config2)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode("1")
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode("0")
			}
		}
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_login_acaunt(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_login_acaunt post")
		var data map[string]interface{}
		json.NewDecoder(r.Body).Decode(&data)

		jsonData, _ := json.Marshal(data)

		resp, _ := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_login, "application/json", bytes.NewBuffer(jsonData))
		defer resp.Body.Close()

		respBody, _ := ioutil.ReadAll(resp.Body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

package register

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// app_back_end/main_com/page/register.go

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

		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		code_ril := GenerateRandomDigits()

		text_xxx := Cripter_xxx(request.Gmail)
		сode := Cripter_xxx(code_ril)

		data := RequestData_xxx{
			Receiver: text_xxx,
			Code:     сode,
		}

		url := config_main.Server_register_and_data_url + config_main.Server_register_and_data_url_send_email
		SendPostRequest_xxx(url, data)

		config := Config_reg{
			Name:   request.Name,
			Pasw:   request.Password,
			Gmail:  request.Gmail,
			Code:   Encrypt_code_reg_save(code_ril),
			Acsses: "0",
		}

		Save_data_reg(config)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_code_verefic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

		code_ril := Decrypt_code_reg_save(user.Code)

		if request.Code == code_ril {
			config := Get_config_user()

			config1 := Config_reg{
				Name:   config.Name,
				Pasw:   config.Pasw,
				Gmail:  config.Gmail,
				Code:   config.Code,
				Acsses: config.Acsses,
			}

			status := Send_user_data_server(config1)

			if status == "1" {
				config2 := Config_reg{
					Name:   config.Name,
					Pasw:   config.Pasw,
					Gmail:  config.Gmail,
					Code:   config.Code,
					Acsses: "1",
				}

				Save_data_reg(config2)

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
		var data map[string]interface{}
		json.NewDecoder(r.Body).Decode(&data)

		jsonData, _ := json.Marshal(data)

		cript_text := Cripter_xxx(string(jsonData))
		data12 := map[string]string{"data": cript_text}
		jsonPayload, _ := json.Marshal(data12)

		resp, _ := http.Post(config_main.Server_register_and_data_url+config_main.Server_register_and_data_url_login, "application/json", bytes.NewBuffer(jsonPayload))
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var result struct {
			Data string `json:"data"`
		}

		if err := json.Unmarshal(body, &result); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		content := Decrypter_AES256(result.Data)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(content)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_reg_file_unix(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data map[string]interface{}

		json.NewDecoder(r.Body).Decode(&data)

		name, _ := data["name"].(string)
		pasw, _ := data["pasw"].(string)
		gmail, _ := data["gmail"].(string)
		code, _ := data["code"].(string)

		config := Config_reg{
			Name:   name,
			Pasw:   pasw,
			Gmail:  gmail,
			Code:   code,
			Acsses: "1",
		}

		Save_data_reg(config)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

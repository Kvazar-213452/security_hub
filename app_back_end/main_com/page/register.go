package page

import (
	"encoding/json"
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

		url := "http://127.0.0.1:5000/send_email"
		page_func.SendPostRequest_xxx(url, data)

		config := page_func.Config_reg{
			Name:  request.Name,
			Pasw:  request.Password,
			Gmail: request.Gmail,
			Code:  page_func.Encrypt_code_reg_save(code_ril),
		}

		page_func.Save_data_reg(config)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
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

		filePath := "../data/user.json"
		fileContent, _ := ioutil.ReadFile(filePath)

		var user User

		json.Unmarshal(fileContent, &user)

		code_ril := page_func.Decrypt_code_reg_save(user.Code)

		if request.Code == code_ril {

		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

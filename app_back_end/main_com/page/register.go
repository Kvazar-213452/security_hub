package page

import (
	"encoding/json"
	"fmt"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"net/http"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_send_email(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/send_email post")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Неможливо прочитати тіло запиту", http.StatusInternalServerError)
			return
		}

		var request struct {
			Name     string `json:"name"`
			Gmail    string `json:"gmail"`
			Password string `json:"password"`
		}

		json.Unmarshal(body, &request)

		text_xxx := "1"
		сode := "ee"

		page_func.Cripter_xxx()

		fmt.Println(сode)
		fmt.Println(text_xxx)

		// data := page_func.RequestData_xxx{
		// 	Receiver: text_xxx,
		// 	Code:     сode,
		// }

		// url := "http://127.0.0.1:5000/send_email"
		// page_func.SendPostRequest_xxx(url, data)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

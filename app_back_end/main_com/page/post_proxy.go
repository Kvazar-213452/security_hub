package page

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"log"
	"net/http"
)

// app_back_end/main_com/page/post_proxy.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_post_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_post_file post")

		proxyReq, err := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_upload, r.Body)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		proxyReq.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_search_server(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_search_server post")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		proxyReq, err := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_search, bytes.NewReader(body))
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for k, v := range r.Header {
			proxyReq.Header[k] = v
		}

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error:")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Post_get_how_many(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_get_how_many post")

		proxyReq, err := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_get_how_many, nil)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		proxyReq.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Error from server:", resp.Status)
			http.Error(w, "Error from server: "+resp.Status, http.StatusInternalServerError)
			return
		}

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println("Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

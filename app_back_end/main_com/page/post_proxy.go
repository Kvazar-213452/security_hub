package page

import (
	"bytes"
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"net/http"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_post_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_post_file post")

		proxyReq, _ := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_upload, r.Body)

		proxyReq.Header = r.Header

		client := &http.Client{}
		resp, _ := client.Do(proxyReq)
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_search_server(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_search_server post")

		body, _ := io.ReadAll(r.Body)

		proxyReq, _ := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_search, bytes.NewReader(body))

		for k, v := range r.Header {
			proxyReq.Header[k] = v
		}

		client := &http.Client{}
		resp, _ := client.Do(proxyReq)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_get_how_many(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_get_how_many post")

		proxyReq, _ := http.NewRequest(http.MethodPost, config_main.Server_data_file_url+config_main.Server_data_file_url_get_how_many, nil)

		proxyReq.Header = r.Header

		client := &http.Client{}
		resp, _ := client.Do(proxyReq)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Error from server: "+resp.Status, http.StatusInternalServerError)
			return
		}

		var response map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

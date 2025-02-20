package main_com

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// auto_update/main_com/post.go

func Post_get_current_ver(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		version := File_config_get_version()

		w.Write([]byte(strconv.Itoa(version)))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Post_get_server_ver(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		version := Get_version()

		w.Write([]byte(strconv.Itoa(version)))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Post_get_url_desc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		url := "http://localhost:5000/get_desc_url"

		data := map[string]string{}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error making POST request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		w.Write([]byte(string(body)))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Post_close(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cmd := exec.Command("../app_back_end/head.exe")
		cmd.Dir = "../app_back_end"

		err := cmd.Start()
		if err != nil {
			http.Error(w, "Error starting head.exe: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("null"))

		fmt.Println("Server and process will now exit...")
		os.Exit(0)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

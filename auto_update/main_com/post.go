package main_com

import (
	"bytes"
	"encoding/json"
	"fmt"
	"head/main_com/update"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("start head.exe...")

	cmd := exec.Command("../app_back_end/head.exe")
	cmd.Dir = "../app_back_end"

	err := cmd.Start()
	if err != nil {
		http.Error(w, "Error starting head.exe: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Помилка запуску head.exe:", err)
		return
	}

	w.Write([]byte("null"))
	fmt.Println("head.exe запущено, сервер закривається...")

	go func() {
		os.Exit(0)
	}()
}

func Post_updata_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		update.DeleteFolders()

		err := update.DownloadAndExtract()
		if err != nil {
			w.Write([]byte("0"))
		} else {
			w.Write([]byte("1"))
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

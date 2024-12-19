package page

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"head/main_com/page_func/background"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

type RequestData struct {
	URL []string `json:"url_site"`
}

type Data struct {
	Value  int    `json:"data"`
	Value1 string `json:"data1"`
}

func Post_antivirus_web(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var requestData RequestData
		json.Unmarshal(body, &requestData)
		url := requestData.URL[0]
		data_good := page_func.CheckUrlInFile(url)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data_good)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_bekend(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		saveDir := "data/bekend"
		os.MkdirAll(saveDir, os.ModePerm)

		func_all.ClearDirectory(saveDir)

		file, fileHeader, _ := r.FormFile("file")
		defer file.Close()

		filePath := filepath.Join(saveDir, fileHeader.Filename)
		destFile, _ := os.Create(filePath)
		defer destFile.Close()

		io.Copy(destFile, file)
		data_good := page_func.Scan_file_virus(filePath)

		json.NewEncoder(w).Encode(data_good)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_change_val_gb_usb(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/change_val_gb_usb change post")

		var data Data
		json.NewDecoder(r.Body).Decode(&data)

		if data.Value == 1 {
			go background.MonitorFlashDrives(config_main.Stop_antivirus_flash_drive, config_main.Antivirus_flash_drive_cmd)
		} else {
			background.Stop_antivirus_flash_drive_func()
		}

		page_func.UpdateConfigKey("antivirus_flash_drive", fmt.Sprintf("%d", data.Value))
		page_func.UpdateConfigKey("antivirus_flash_drive_cmd", fmt.Sprintf("%s", data.Value1))

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

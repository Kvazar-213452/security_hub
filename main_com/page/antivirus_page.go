package page

import (
	"encoding/json"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
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

func Post_antivirus_web(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не вдалося прочитати тіло запиту", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var requestData RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil || len(requestData.URL) == 0 {
			http.Error(w, "Некоректний формат запиту", http.StatusBadRequest)
			return
		}

		url := requestData.URL[0]
		page_func.FetchHTMLAndJS(url)
		code := page_func.CheckUrlInFile(url)

		jsonData := func_all.ReadFileToJSON("data/inter.txt")

		func_all.Clear_file(config_main.Global_phat + "\\data\\inter.txt")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if code == 1 {
			response := map[string]interface{}{
				"found": true,
				"data":  jsonData,
			}
			json.NewEncoder(w).Encode(response)
		} else {
			response := map[string]interface{}{
				"found": false,
				"data":  jsonData,
			}
			json.NewEncoder(w).Encode(response)
		}

		page_func.DeleteFiles()
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_bekend(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		saveDir := "data/bekend"
		os.MkdirAll(saveDir, os.ModePerm)

		func_all.ClearDirectory(saveDir)

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Не вдалося отримати файл", http.StatusBadRequest)
			return
		}
		defer file.Close()

		filePath := filepath.Join(saveDir, fileHeader.Filename)
		destFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Не вдалося створити файл на сервері", http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, file)
		if err != nil {
			http.Error(w, "Не вдалося зберегти файл", http.StatusInternalServerError)
			return
		}

		value := r.FormValue("value")

		if value == "0" {
			data := page_func.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_2)

			if data == 0 {
				w.Write([]byte("0"))
			} else {
				w.Write([]byte("1"))
			}
		} else {
			data := page_func.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_1)
			data1 := page_func.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_2)

			if data == 0 || data1 == 1 {
				w.Write([]byte("0"))
			} else {
				w.Write([]byte("1"))
			}
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

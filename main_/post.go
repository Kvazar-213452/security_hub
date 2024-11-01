package main_

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"head/main_/antivirus"
	config_main "head/main_/config"
	"head/main_/encryption"
	"head/main_/func_all"
	page_func "head/main_/page_func_spec"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type VisualizationMessage struct {
	Message int `json:"message"`
}

type RequestData struct {
	URL []string `json:"url_site"`
}

type OSData struct {
	SystemMemory  string `json:"system_memory"`
	ProcessorInfo string `json:"processor_info"`
	OSVersion     string `json:"os_version"`
	ComputerName  string `json:"computer_name"`
	UserName      string `json:"user_name"`
	SystemUptime  string `json:"system_uptime"`
}

func Post_gagat_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_wifi_now post")

		wifiInfo, err := func_all.Get_Wifi_info()
		if err != nil {
			http.Error(w, "Помилка отримання інформації про Wi-Fi", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wifiInfo)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_wifi_network(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_wifi post")
		networks, err := func_all.Get_available_Wifi_networks()
		if err != nil {
			json.NewEncoder(w).Encode("error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(networks)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_server_fet_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/get_logs post")

		jsonData, err := func_all.LoadLogFile()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_network_now(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/network_now post")

		ssid := func_all.GetConnectedSSID()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"ssid": ssid}); err != nil {
			http.Error(w, "Помилка при кодуванні JSON", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_config_global(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/config_global post")

		config, err := func_all.LoadConfig()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(config)
		if err != nil {
			http.Error(w, "не вдалося закодувати в JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_config_change(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/visualization change post")

		var msg VisualizationMessage

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Не вдалося декодувати JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		func_all.UpdateVisualization(strconv.Itoa(msg.Message), "Visualization")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_get_os_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("get os data")

		systemMemory := page_func.GetSystemMemory()
		processorInfo := page_func.GetProcessorInfo()
		osVersion := page_func.GetOSVersion()
		computerName := page_func.GetComputerNameCustom()
		userName := page_func.GetUserNameCustom()
		systemUptime := page_func.GetSystemUptime()

		osData := OSData{
			SystemMemory:  systemMemory,
			ProcessorInfo: processorInfo,
			OSVersion:     osVersion,
			ComputerName:  computerName,
			UserName:      userName,
			SystemUptime:  systemUptime,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(osData)
		if err != nil {
			http.Error(w, "Помилка при формуванні JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_usb_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("usb info")

		info := func_all.Usb_info()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"devices": devices}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Не вдалося кодувати JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_resource_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("resource info")

		info := page_func.Get_all_data_now()
		cleanedInfo := strings.ReplaceAll(info, "\r", "")
		devices := strings.Split(cleanedInfo, "\n")
		response := map[string][]string{"data": devices}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Не вдалося кодувати JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_cleanup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("cleanup")

		func_all.Cleanup()

		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
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
		antivirus.FetchHTMLAndJS(url)
		code := antivirus.CheckUrlInFile(url)

		jsonData := func_all.ReadFileToJSON("data/inter.txt")

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

		antivirus.DeleteFiles()
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_bekend(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		saveDir := "data/bekend"
		err := os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Не вдалося створити директорію", http.StatusInternalServerError)
			return
		}

		err = func_all.ClearDirectory(saveDir)
		if err != nil {
			http.Error(w, "Не вдалося очистити директорію", http.StatusInternalServerError)
			return
		}

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
			data := antivirus.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_2)

			if data == 0 {
				w.Write([]byte("0"))
			} else {
				w.Write([]byte("1"))
			}
		} else {
			data := antivirus.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_1)
			data1 := antivirus.Scan_file_virus(filePath, config_main.Server_data_sha1_hashes_2)

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

func Post_encryption_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("encryption_file")

		filename_u := r.FormValue("filename")
		filename := "data/encryption/" + filename_u

		err := func_all.ClearDirectory("data/encryption")
		if err != nil {
			http.Error(w, "Не вдалося очистити директорію", http.StatusInternalServerError)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Помилка при читанні файлу", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		savePath := "data/encryption"
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			http.Error(w, "Помилка при створенні директорії", http.StatusInternalServerError)
			return
		}

		filePath := filepath.Join(savePath, header.Filename)

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Помилка при створенні файлу", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Помилка при збереженні файлу", http.StatusInternalServerError)
			return
		}

		key := encryption.GenerateKey()

		encryptedContent, err := encryption.EncryptFile(filename, key)
		if err != nil {
			w.Write([]byte("0"))
			return
		}

		encFilePath := config_main.Frontend_folder + "/static/data/main.enc"
		err = os.WriteFile(encFilePath, encryptedContent, 0644)
		if err != nil {
			http.Error(w, "Помилка при збереженні зашифрованого файлу", http.StatusInternalServerError)
			return
		}

		keyHex := hex.EncodeToString(key)
		w.Write([]byte(keyHex))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

func Post_decipher_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("decipher_file")

		key := r.FormValue("key")
		filename := "data/decipher/" + "main.enc"

		func_all.ClearDirectory("data/decipher")

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Помилка при читанні файлу", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		savePath := "data/decipher"
		os.MkdirAll(savePath, os.ModePerm)

		filePath := filepath.Join(savePath, "main.enc")

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Помилка при створенні файлу", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Помилка при збереженні файлу", http.StatusInternalServerError)
			return
		}

		err = encryption.DecryptFile(filename, key)
		if err != nil {
			fmt.Println("Помилка:", err)
			w.Write([]byte("0"))
			return
		}

		w.Write([]byte("1"))
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

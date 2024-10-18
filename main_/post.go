package main_

import (
	"encoding/json"
	"head/main_/antivirus"
	"head/main_/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
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
			http.Error(w, "Помилка отримання інформації про Wi-Fi мережі", http.StatusInternalServerError)
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

		dll := syscall.NewLazyDLL("library/system_info.dll")

		getSystemMemory := dll.NewProc("GetSystemMemory")
		getProcessorInfo := dll.NewProc("GetProcessorInfo")
		getOSVersion := dll.NewProc("GetOSVersion")
		getComputerNameCustom := dll.NewProc("GetComputerNameCustom")
		getUserNameCustom := dll.NewProc("GetUserNameCustom")
		getSystemUptime := dll.NewProc("GetSystemUptime")

		systemMemory := func_all.RemoveNewlines(func_all.CallDLLFunction(getSystemMemory, "System Memory Info"))
		processorInfo := func_all.RemoveNewlines(func_all.CallDLLFunction(getProcessorInfo, "Processor Info"))
		osVersion := func_all.RemoveNewlines(func_all.CallDLLFunction(getOSVersion, "OS Version Info"))
		computerName := func_all.RemoveNewlines(func_all.CallDLLFunction(getComputerNameCustom, "Computer Name Info"))
		userName := func_all.RemoveNewlines(func_all.CallDLLFunction(getUserNameCustom, "User Name Info"))
		systemUptime := func_all.RemoveNewlines(func_all.CallDLLFunction(getSystemUptime, "System Uptime Info"))

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

		info := func_all.Resource_info()
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
			data := antivirus.Scan_file_virus(filePath, "data/sha1_hashes_2.txt")

			if data == 0 {
				w.Write([]byte("0"))
			} else {
				w.Write([]byte("1"))
			}
		} else {
			data := antivirus.Scan_file_virus(filePath, "data/sha1_hashes_1.txt")
			data1 := antivirus.Scan_file_virus(filePath, "data/sha1_hashes_2.txt")

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

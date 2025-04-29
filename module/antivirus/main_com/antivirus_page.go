package main_com

import (
	"encoding/json"
	"head/main_com/antivirus"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// app_back_end/main_com/page/antivirus_page.go

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
		data_good := antivirus.CheckUrlInFile(url)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data_good)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
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
		data_good := antivirus.Scan_file_virus(filePath)

		json.NewEncoder(w).Encode(data_good)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_bekend_scan_dir(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_antivirus_bekend_scan_dir")

		body, _ := io.ReadAll(r.Body)

		var request struct {
			Dir string `json:"dir"`
		}
		json.Unmarshal(body, &request)

		exeFiles := antivirus.Scan_exeFiles(request.Dir)

		resultData := map[string]interface{}{
			"total_exe_files":  len(exeFiles),
			"checked_files":    []map[string]string{},
			"detected_viruses": []map[string]string{},
		}

		for _, exeFile := range exeFiles {
			fileHash := antivirus.File_hash(exeFile)

			result := antivirus.Check_hash_VirusTotal(fileHash)

			checkedFile := map[string]string{
				"path": exeFile,
				"hash": fileHash,
			}
			resultData["checked_files"] = append(resultData["checked_files"].([]map[string]string), checkedFile)

			if result != nil {
				if data, ok := result["data"].(map[string]interface{}); ok {
					if attributes, ok := data["attributes"].(map[string]interface{}); ok {
						if lastAnalysisStats, ok := attributes["last_analysis_stats"].(map[string]interface{}); ok {
							if malicious, ok := lastAnalysisStats["malicious"].(float64); ok && malicious > 0 {
								detectedVirus := map[string]string{
									"path": exeFile,
									"hash": fileHash,
								}
								resultData["detected_viruses"] = append(resultData["detected_viruses"].([]map[string]string), detectedVirus)
							}
						}
					}
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resultData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_bekend_del_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_antivirus_bekend_scan_dir")

		body, _ := io.ReadAll(r.Body)

		var request struct {
			Path string `json:"path"`
		}
		json.Unmarshal(body, &request)

		status := antivirus.Delete_file(request.Path)

		resultData := map[string]interface{}{
			"status": status,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resultData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_antivirus_resurse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("/Post_antivirus_resurse")

		data := antivirus.Get_process_info()

		resultData := map[string]interface{}{
			"status": data,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resultData)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

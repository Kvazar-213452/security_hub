package main_com

import (
	"encoding/json"
	"flash_drive/main_com/antivirus"
	"flash_drive/main_com/func_all"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

func StartServer(port int) {
	portStr := ":" + strconv.Itoa(port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			os.Exit(0)
		} else {
			http.Error(w, "error", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/del_file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error body: %v", err), http.StatusBadRequest)
				return
			}

			var requestData map[string]interface{}
			if err := json.Unmarshal(body, &requestData); err != nil {
				http.Error(w, fmt.Sprintf("Error JSON: %v", err), http.StatusBadRequest)
				return
			}

			os.Remove(requestData["data"].(string))

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("ss")
		} else {
			http.Error(w, "error", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/get_file", func(w http.ResponseWriter, r *http.Request) {
		exePath, err := os.Executable()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting executable path: %v", err), http.StatusInternalServerError)
			return
		}
		exeDir := filepath.Dir(exePath)

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error body: %v", err), http.StatusBadRequest)
			return
		}

		var requestData map[string]interface{}
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, fmt.Sprintf("Error JSON: %v", err), http.StatusBadRequest)
			return
		}

		data, ok := requestData["data"].(string)
		if !ok {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		fullPath := filepath.Join(exeDir, data)
		fullPath = filepath.Clean(fullPath)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	})

	http.HandleFunc("/scan_flash", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		flashLetter, err := antivirus.ReadFlashDriveLetter("flash.md")
		if err != nil {
			http.Error(w, "Cannot read flash.md: "+err.Error(), http.StatusInternalServerError)
			return
		}

		exeFiles, err := antivirus.FindExeFiles(flashLetter)
		if err != nil {
			http.Error(w, "Error scanning files: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var maliciousFiles []string
		for _, file := range exeFiles {
			hash, err := antivirus.CalculateSHA256(file)
			if err != nil {
				continue
			}
			isMalicious, err := antivirus.CheckVirusTotalHash(hash)
			if err != nil {
				continue
			}
			if isMalicious {
				maliciousFiles = append(maliciousFiles, file)
			}
		}

		result := antivirus.ScanResult{
			AllExeFiles:    exeFiles,
			MaliciousFiles: maliciousFiles,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	// page

	http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"web/scan.html",
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "scan.html", nil)
	})

	func_all.WriteServerInfo(portStr)

	if err := http.ListenAndServe(portStr, nil); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

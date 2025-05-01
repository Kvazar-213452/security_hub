package main_com

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/encryption"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// app_back_end/main_com/page/encryption_page.go

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_encryption_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		filename_u := r.FormValue("filename")
		filename := "data/encryption/" + filename_u

		func_all.ClearDirectory("data/encryption")
		func_all.ClearDirectory("web/static/data")

		file, header, _ := r.FormFile("file")
		defer file.Close()

		savePath := "data/encryption"
		os.MkdirAll(savePath, os.ModePerm)

		filePath := filepath.Join(savePath, header.Filename)

		dst, _ := os.Create(filePath)
		defer dst.Close()

		io.Copy(dst, file)
		key := encryption.GenerateKey()

		encryptedContent, err := encryption.EncryptFile(filename, key)
		if err != nil {
			w.Write([]byte("0"))
			return
		}

		encFilePath := config_main.Frontend_folder + "/static/data/main.enc"
		os.WriteFile(encFilePath, encryptedContent, 0644)

		keyHex := hex.EncodeToString(key)
		w.Write([]byte(keyHex))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_decipher_file(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		key := r.FormValue("key")
		filename := "data/decipher/" + "main.enc"

		func_all.ClearDirectory("data/decipher")
		func_all.ClearDirectory("web/static/data")

		file, _, _ := r.FormFile("file")
		defer file.Close()

		savePath := "data/decipher"
		os.MkdirAll(savePath, os.ModePerm)

		filePath := filepath.Join(savePath, "main.enc")

		dst, _ := os.Create(filePath)
		defer dst.Close()

		io.Copy(dst, file)

		err := encryption.DecryptFile(filename, key)
		if err != nil {
			fmt.Println("Помилка:", err)
			w.Write([]byte("0"))
			return
		}

		w.Write([]byte("1"))
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Get_file(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	var requestData map[string]interface{}
	json.Unmarshal(body, &requestData)

	filePath := filepath.Join(requestData["data"].(string))

	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

package page

import (
	"encoding/hex"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page_func"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

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

		key := page_func.GenerateKey()

		encryptedContent, err := page_func.EncryptFile(filename, key)
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

		err = page_func.DecryptFile(filename, key)
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

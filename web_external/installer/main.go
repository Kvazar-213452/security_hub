package main

import (
	"fmt"
	"head/main_com"
	"net/http"
)

func main() {
	defer fmt.Println("HHH")

	base64Data := main_com.Base64_code
	outputFilePath := "output_file.zip"

	err := main_com.Decode_Base64_ToFile(base64Data, outputFilePath)
	if err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Println("Файл успішно створено:", outputFilePath)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, main_com.Page_1)
	})
}

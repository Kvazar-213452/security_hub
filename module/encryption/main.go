package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"os"
	"strconv"
)

// module/encryption/main.go

func main() {
	var port int

	if p, err := strconv.Atoi(os.Args[1]); err == nil {
		port = p
	} else {
		os.Exit(1)
	}

	portStr := ":" + strconv.Itoa(port)
	func_all.Config_port(strconv.Itoa(port))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/", main_com.Render_encryption)

	http.HandleFunc("/encryption_file", main_com.Post_encryption_file)
	http.HandleFunc("/decipher_file", main_com.Post_decipher_file)
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

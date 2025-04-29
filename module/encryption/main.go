package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"strconv"
)

func main() {
	port := 4300 //func_all.FindFreePort()
	portStr := ":" + strconv.Itoa(port)
	func_all.Config_port(strconv.Itoa(port))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/main", main_com.Render_encryption)

	http.HandleFunc("/encryption_file", main_com.Post_encryption_file)
	http.HandleFunc("/decipher_file", main_com.Post_decipher_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

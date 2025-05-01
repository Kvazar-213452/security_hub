package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"os"
	"strconv"
)

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

	http.HandleFunc("/", main_com.Render_password_page)

	http.HandleFunc("/get_password", main_com.Post_get_password)
	http.HandleFunc("/add_key_pasw", main_com.Post_add_key_pasw)
	http.HandleFunc("/del_key_pasw", main_com.Post_del_key_pasw)
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

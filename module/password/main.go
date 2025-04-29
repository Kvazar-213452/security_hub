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

	http.HandleFunc("/main", main_com.Render_password_page)

	http.HandleFunc("/get_password", main_com.Post_get_password)
	http.HandleFunc("/add_key_pasw", main_com.Post_add_key_pasw)
	http.HandleFunc("/del_key_pasw", main_com.Post_del_key_pasw)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

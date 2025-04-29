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

	http.HandleFunc("/main", main_com.Render_system_page)

	http.HandleFunc("/get_os_data", main_com.Post_get_os_data)
	http.HandleFunc("/window_open", main_com.Post_window_open)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

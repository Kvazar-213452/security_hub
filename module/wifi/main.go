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

	http.HandleFunc("/main", main_com.Render_wifi_page)

	http.HandleFunc("/get_wifi_now", main_com.Post_gagat_network)
	http.HandleFunc("/get_wifi", main_com.Post_wifi_network)
	http.HandleFunc("/get_pacage_info", main_com.Post_get_pacage_info_wifi)
	http.HandleFunc("/network_now", main_com.Post_network_now)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

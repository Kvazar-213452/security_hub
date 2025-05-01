package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"os"
	"strconv"
)

// module/wifi/main.go

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

	http.HandleFunc("/main", main_com.Render_wifi_page)
	http.HandleFunc("/get_wifi_now", main_com.Post_gagat_network)
	http.HandleFunc("/get_wifi", main_com.Post_wifi_network)
	http.HandleFunc("/get_pacage_info", main_com.Post_get_pacage_info_wifi)
	http.HandleFunc("/network_now", main_com.Post_network_now)
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		os.Exit(1)
	}
}

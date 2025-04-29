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

	http.HandleFunc("/main", main_com.Render_antivirus)

	http.HandleFunc("/antivirus_web", main_com.Post_antivirus_web)
	http.HandleFunc("/antivirus_bekend", main_com.Post_antivirus_bekend)
	http.HandleFunc("/antivirus_resurse", main_com.Post_antivirus_resurse)
	http.HandleFunc("/antivirus_bekend_scan_dir", main_com.Post_antivirus_bekend_scan_dir)
	http.HandleFunc("/antivirus_bekend_del_file", main_com.Post_antivirus_bekend_del_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

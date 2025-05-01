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
	func_all.Config_port(portStr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/", main_com.Render_antivirus)

	http.HandleFunc("/antivirus_web", main_com.Post_antivirus_web)
	http.HandleFunc("/antivirus_bekend", main_com.Post_antivirus_bekend)
	http.HandleFunc("/antivirus_resurse", main_com.Post_antivirus_resurse)
	http.HandleFunc("/antivirus_bekend_scan_dir", main_com.Post_antivirus_bekend_scan_dir)
	http.HandleFunc("/antivirus_bekend_del_file", main_com.Post_antivirus_bekend_del_file)
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

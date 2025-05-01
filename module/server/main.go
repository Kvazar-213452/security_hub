package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"os"
	"strconv"
)

// module/server/main.go

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

	http.HandleFunc("/", main_com.Render_server_page)

	// post
	http.HandleFunc("/search_server", main_com.Post_search_server)
	http.HandleFunc("/post_file_server", main_com.Post_post_file)
	http.HandleFunc("/get_how_many", main_com.Post_get_how_many)
	http.HandleFunc("/browser_site_server", main_com.Post_Browser_site_server)
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

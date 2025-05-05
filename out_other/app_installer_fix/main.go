package main

import (
	"fmt"
	"head/main_com"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/browser"
)

func main() {
	port := 4000 // func_all.FindFreePort()

	portStr := ":" + strconv.Itoa(port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/", main_com.Render_main_page)
	http.HandleFunc("/off", main_com.Get_off_app)

	// post

	http.HandleFunc("/install", main_com.Post_install)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	err := http.ListenAndServe(portStr, nil)
	browser.OpenURL("http://localhost" + portStr)
	if err != nil {
		os.Exit(1)
	}
}

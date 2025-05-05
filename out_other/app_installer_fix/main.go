package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/func_all"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/browser"
)

func main() {
	port := func_all.FindFreePort()
	portStr := ":" + strconv.Itoa(port)
	url := "http://localhost" + portStr

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", main_com.Render_main_page)
	http.HandleFunc("/off", main_com.Get_off_app)
	http.HandleFunc("/install", main_com.Post_install)

	go func() {
		fmt.Println("server on", url)
		err := http.ListenAndServe(portStr, nil)
		if err != nil {
			fmt.Println("error server:", err)
			os.Exit(1)
		}
	}()

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Println("error", err)
	}

	select {}
}

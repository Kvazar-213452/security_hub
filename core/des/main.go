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
	err := main_com.RunModules(
		"../data/config_module.json",
		"result.json",
	)
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
		os.Exit(1)
	}

	port := 4300
	portStr := ":" + strconv.Itoa(port)
	func_all.Starter(strconv.Itoa(port))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/register", main_com.Render_register_page)
	http.HandleFunc("/settings", main_com.Render_settings_page)

	// post
	http.HandleFunc("/get_file", main_com.Get_file)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

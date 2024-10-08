package main

import (
	"fmt"
	"head/main_"
	"net/http"
)

func main() {
	config := main_.LoadConfig("config.toml")
	port := fmt.Sprintf(":%d", config.Port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/save_file", main_.Render_index_page)

	fmt.Printf("Сервер працює на порту %d\n", config.Port)
	http.ListenAndServe(port, nil)
}

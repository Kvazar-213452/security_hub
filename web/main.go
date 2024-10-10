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

	// page
	http.HandleFunc("/", main_.Render_index_page)
	http.HandleFunc("/about", main_.Render_about_page)
	http.HandleFunc("/settings", main_.Render_settings_page)

	// post
	http.HandleFunc("/endpoint", main_.Post_gagat_network)

	fmt.Printf("Сервер працює на порту %d\n", config.Port)
	http.ListenAndServe(port, nil)
}

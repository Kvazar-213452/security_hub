package main

import (
	"fmt"
	"head/main_"
	"net/http"
	"strconv"
)

func main() {
	port := main_.FindFreePort()
	portStr := ":" + strconv.Itoa(port)

	main_.Write_config_core(portStr)
	cmd := main_.StartShellWeb()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Page
	http.HandleFunc("/", main_.Render_index_page)
	http.HandleFunc("/about", main_.Render_about_page)
	http.HandleFunc("/settings", main_.Render_settings_page)
	http.HandleFunc("/system", main_.Render_system_page)

	// Post
	http.HandleFunc("/get_wifi_now", main_.Post_gagat_network)
	http.HandleFunc("/get_wifi", main_.Post_wifi_network)
	http.HandleFunc("/get_logs", main_.Post_server_fet_log)
	http.HandleFunc("/network_now", main_.Post_network_now)
	http.HandleFunc("/config_global", main_.Post_config_global)
	http.HandleFunc("/visualization", main_.Post_config_change)

	fmt.Printf("Сервер працює на порту %d\n", port)
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		fmt.Printf("Помилка запуску сервера: %v\n", err)
	}

	if err := cmd.Process.Kill(); err != nil {
		fmt.Printf("Не вдалося завершити shell_web.exe: %v\n", err)
	}
}

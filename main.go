package main

import (
	"fmt"
	"head/main_"
	config_main "head/main_/config"
	"head/main_/func_all"
	"net/http"
	"os/exec"
	"strconv"
)

func main() {
	config, err := func_all.LoadConfig_start(config_main.Main_config)
	if err != nil {
		fmt.Printf("Не вдалося завантажити конфігурацію: %v\n", err)
		return
	}

	var port int
	if config.Port == 0 {
		port = func_all.FindFreePort()
	} else {
		port = config.Port
	}

	portStr := ":" + strconv.Itoa(port)

	var cmd *exec.Cmd
	if config.Visualization == 1 {
		func_all.Write_config_core(portStr)
		cmd = func_all.StartShellWeb()
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("front_end/static"))))

	// get
	http.HandleFunc("/", main_.Render_index_page)
	http.HandleFunc("/about", main_.Render_about_page)
	http.HandleFunc("/settings", main_.Render_settings_page)
	http.HandleFunc("/system", main_.Render_system_page)
	http.HandleFunc("/off_app", main_.Post_off_app)

	// Post
	http.HandleFunc("/get_wifi_now", main_.Post_gagat_network)
	http.HandleFunc("/get_wifi", main_.Post_wifi_network)
	http.HandleFunc("/get_logs", main_.Post_server_fet_log)
	http.HandleFunc("/network_now", main_.Post_network_now)
	http.HandleFunc("/config_global", main_.Post_config_global)
	http.HandleFunc("/visualization", main_.Post_config_change)
	http.HandleFunc("/get_os_data", main_.Post_get_os_data)
	http.HandleFunc("/usb_info", main_.Post_usb_info)
	http.HandleFunc("/resource_info", main_.Post_resource_info)

	fmt.Printf("Сервер працює на порту %d\n", port)
	err = http.ListenAndServe(portStr, nil)
	if err != nil {
		fmt.Printf("Помилка запуску сервера: %v\n", err)
	}

	if cmd != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Printf("Не вдалося завершити shell_web.exe: %v\n", err)
		}
	}
}

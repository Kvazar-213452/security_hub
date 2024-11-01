package main

import (
	"fmt"
	"head/main_"
	config_main "head/main_/config"
	"head/main_/func_all"
	"head/main_/page"
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
		cmd = func_all.StartShellWeb(portStr)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("front_end/static"))))

	// get
	http.HandleFunc("/wifi", main_.Render_index_page)
	http.HandleFunc("/about", main_.Render_about_page)
	http.HandleFunc("/settings", main_.Render_settings_page)
	http.HandleFunc("/system", main_.Render_system_page)
	http.HandleFunc("/off_app", main_.Get_off_app)
	http.HandleFunc("/cleaning", main_.Render_cleaning)
	http.HandleFunc("/antivirus", main_.Render_antivirus)
	http.HandleFunc("/encryption", main_.Render_encryption)
	http.HandleFunc("/browser_site_app", main_.Browser_site_app)

	// Post
	http.HandleFunc("/get_wifi_now", page.Post_gagat_network)
	http.HandleFunc("/get_wifi", page.Post_wifi_network)
	http.HandleFunc("/get_logs", main_.Post_server_fet_log)
	http.HandleFunc("/network_now", page.Post_network_now)
	http.HandleFunc("/config_global", page.Post_config_global)
	http.HandleFunc("/visualization", page.Post_config_change)
	http.HandleFunc("/get_os_data", page.Post_get_os_data)
	http.HandleFunc("/usb_info", page.Post_usb_info)
	http.HandleFunc("/resource_info", page.Post_resource_info)
	http.HandleFunc("/cleanup", page.Post_cleanup)
	http.HandleFunc("/antivirus_web", page.Post_antivirus_web)
	http.HandleFunc("/antivirus_bekend", page.Post_antivirus_bekend)
	http.HandleFunc("/encryption_file", page.Post_encryption_file)
	http.HandleFunc("/decipher_file", page.Post_decipher_file)

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

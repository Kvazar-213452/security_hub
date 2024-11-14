package main

import (
	"fmt"
	"head/main_com"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"head/main_com/page"
	"head/main_com/page_func/background"
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

	if config.Antivirus.Antivirus_flash_drive == 1 {
		go background.MonitorFlashDrives(config_main.Stop_antivirus_flash_drive, config_main.Antivirus_flash_drive_cmd)
	}

	portStr := ":" + strconv.Itoa(port)

	var cmd *exec.Cmd
	if config.Visualization == 1 {
		cmd = func_all.StartShellWeb(port, config.Shell)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// get
	http.HandleFunc("/main", main_com.Render_index_page)
	http.HandleFunc("/about", main_com.Render_about_page)
	http.HandleFunc("/settings", main_com.Render_settings_page)
	http.HandleFunc("/system", main_com.Render_system_page)
	http.HandleFunc("/off_app", main_com.Get_off_app)
	http.HandleFunc("/cleaning", main_com.Render_cleaning)
	http.HandleFunc("/antivirus", main_com.Render_antivirus)
	http.HandleFunc("/encryption", main_com.Render_encryption)
	http.HandleFunc("/wifi", main_com.Render_wifi_page)

	// Post
	http.HandleFunc("/get_wifi_now", page.Post_gagat_network)
	http.HandleFunc("/get_wifi", page.Post_wifi_network)
	http.HandleFunc("/get_logs", main_com.Post_server_fet_log)
	http.HandleFunc("/network_now", page.Post_network_now)
	http.HandleFunc("/config_global", page.Post_config_global)
	http.HandleFunc("/visualization", page.Post_config_change)
	http.HandleFunc("/get_os_data", page.Post_get_os_data)
	http.HandleFunc("/window_open", page.Post_window_open)
	http.HandleFunc("/resource_info", page.Post_resource_info)
	http.HandleFunc("/cleanup", page.Post_cleanup)
	http.HandleFunc("/antivirus_web", page.Post_antivirus_web)
	http.HandleFunc("/antivirus_bekend", page.Post_antivirus_bekend)
	http.HandleFunc("/encryption_file", page.Post_encryption_file)
	http.HandleFunc("/decipher_file", page.Post_decipher_file)
	http.HandleFunc("/log_change", page.Post_log_change)
	http.HandleFunc("/port_change", page.Post_port_change)
	http.HandleFunc("/shell_change", page.Post_shell_change)
	http.HandleFunc("/browser_site_app", main_com.Post_Browser_site_app)
	http.HandleFunc("/change_val_gb_usb", page.Post_change_val_gb_usb)
	http.HandleFunc("/change_lang_settings", page.Post_change_lang_settings)

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

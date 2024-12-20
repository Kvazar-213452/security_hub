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

// ⠄⠄⠄⠄⢠⣿⣿⣿⣿⣿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣯⢻⣿⣿⣿⣿⣆⠄⠄⠄
// ⠄⠄⣼⢀⣿⣿⣿⣿⣏⡏⠄⠹⣿⣿⣿⣿⣿⣿⣿⣿⣧⢻⣿⣿⣿⣿⡆⠄⠄
// ⠄⠄⡟⣼⣿⣿⣿⣿⣿⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⣿⣇⢻⣿⣿⣿⣿⠄⠄
// ⠄⢰⠃⣿⣿⠿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠙⠿⣿⣿⣿⣿⣿⠄⢿⣿⣿⣿⡄⠄
// ⠄⢸⢠⣿⣿⣧⡙⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠈⠛⢿⣿⣿⡇⠸⣿⡿⣸⡇⠄
// ⠄⠈⡆⣿⣿⣿⣿⣦⡙⠳⠄⠄⠄⠄⠄⠄⢀⣠⣤⣀⣈⠙⠃⠄⠿⢇⣿⡇⠄
// ⠄⠄⡇⢿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⣠⣶⣿⣿⣿⣿⣿⣿⣷⣆⡀⣼⣿⡇⠄
// ⠄⠄⢹⡘⣿⣿⣿⢿⣷⡀⠄⢀⣴⣾⣟⠉⠉⠉⠉⣽⣿⣿⣿⣿⠇⢹⣿⠃⠄
// ⠄⠄⠄⢷⡘⢿⣿⣎⢻⣷⠰⣿⣿⣿⣿⣦⣀⣀⣴⣿⣿⣿⠟⢫⡾⢸⡟⠄.
// ⠄⠄⠄⠄⠻⣦⡙⠿⣧⠙⢷⠙⠻⠿⢿⡿⠿⠿⠛⠋⠉⠄⠂⠘⠁⠞⠄⠄⠄
// ⠄⠄⠄⠄⠄⠈⠙⠑⣠⣤⣴⡖⠄⠿⣋⣉⣉⡁⠄⢾⣦⠄⠄⠄⠄⠄⠄⠄⠄

func main() {
	config := func_all.LoadConfig_start(config_main.Main_config)

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
	func_all.Config_port(strconv.Itoa(port))

	var cmd *exec.Cmd
	if config.Visualization == 1 {
		cmd = func_all.StartShellWeb(port, config.Shell, config.Version)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../app_front_end/static"))))

	// get
	http.HandleFunc("/main", main_com.Render_index_page)
	http.HandleFunc("/settings", main_com.Render_settings_page)
	http.HandleFunc("/system", main_com.Render_system_page)
	http.HandleFunc("/off_app", main_com.Get_off_app)
	http.HandleFunc("/cleaning", main_com.Render_cleaning)
	http.HandleFunc("/antivirus", main_com.Render_antivirus)
	http.HandleFunc("/encryption", main_com.Render_encryption)
	http.HandleFunc("/wifi", main_com.Render_wifi_page)
	http.HandleFunc("/server", main_com.Render_server_page)
	http.HandleFunc("/password", main_com.Render_password_page)
	http.HandleFunc("/register", main_com.Render_register_page)

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
	http.HandleFunc("/style_change", page.Post_style_change)
	http.HandleFunc("/get_style", main_com.Post_get_style)
	http.HandleFunc("/install_style", main_com.Post_install_style)
	http.HandleFunc("/scan_dir", page.Post_scan_dir)
	http.HandleFunc("/version_get", main_com.Post_version_get)
	http.HandleFunc("/send_email", page.Post_send_email)
	http.HandleFunc("/code_verefic", page.Post_code_verefic)
	http.HandleFunc("/get_pacage_info", page.Post_get_pacage_info_wifi)
	http.HandleFunc("/post_file_server", page.Post_post_file)
	http.HandleFunc("/search_server", page.Post_search_server)
	http.HandleFunc("/version_get_server", main_com.Post_version_get_server)
	http.HandleFunc("/reg_status", main_com.Post_reg_status)
	http.HandleFunc("/get_password", page.Post_get_password)
	http.HandleFunc("/add_key_pasw", page.Post_add_key_pasw)
	http.HandleFunc("/login_acaunt", page.Post_login_acaunt)
	http.HandleFunc("/del_key_pasw", page.Post_del_key_pasw)

	fmt.Printf("started %d\n", port)
	http.ListenAndServe(portStr, nil)

	if cmd != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Printf("not shell_web.exe: %v\n", err)
		}
	}
}

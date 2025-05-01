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
	var port int

	if p, err := strconv.Atoi(os.Args[1]); err == nil {
		port = p
	} else {
		os.Exit(1)
	}

	err := main_com.RunModules(
		"../data/config_module.json",
		"result.json",
	)
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
		os.Exit(1)
	}

	portStr := ":" + strconv.Itoa(port)
	func_all.Starter(strconv.Itoa(port))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/register", main_com.Render_register_page)
	http.HandleFunc("/settings", main_com.Render_settings_page)
	http.HandleFunc("/off_app", main_com.Get_off_app)

	// post
	http.HandleFunc("/browser_site_app", main_com.Post_Browser_site_app)
	http.HandleFunc("/get_file", main_com.Get_file)
	http.HandleFunc("/config_global", main_com.Post_config_global)
	http.HandleFunc("/visualization", main_com.Post_config_change)
	http.HandleFunc("/log_change", main_com.Post_log_change)
	http.HandleFunc("/port_change", main_com.Post_port_change)
	http.HandleFunc("/shell_change", main_com.Post_shell_change)
	http.HandleFunc("/change_lang_settings", main_com.Post_change_lang_settings)
	http.HandleFunc("/style_change", main_com.Post_style_change)
	http.HandleFunc("/send_email", main_com.Post_send_email)
	http.HandleFunc("/code_verefic", main_com.Post_code_verefic)
	http.HandleFunc("/reg_file_unix", main_com.Post_reg_file_unix)
	http.HandleFunc("/updata_app", main_com.Post_updata_app)
	http.HandleFunc("/accses_updata", main_com.Post_accses_updata)
	http.HandleFunc("/info_module_nm", main_com.Post_info_module_nm)
	http.HandleFunc("/install_module", main_com.Post_install_module)
	http.HandleFunc("/uninstall_module", main_com.Post_uninstall_module)
	http.HandleFunc("/install_style", main_com.Post_install_style)
	http.HandleFunc("/del_temp", main_com.Post_del_temp)
	http.HandleFunc("/get_temp_info", main_com.Post_get_temp_info)
	http.HandleFunc("/version_get_server", main_com.Post_version_get_server)
	http.HandleFunc("/version_get", main_com.Post_version_get)
	http.HandleFunc("/get_info_work_server_register", main_com.Post_get_info_work_server_register)
	http.HandleFunc("/get_info_work_server_data_file", main_com.Post_get_info_work_server_data_file)
	http.HandleFunc("/log_out", main_com.Post_log_out)

	fmt.Println("Сервер запущено на http://localhost" + portStr)
	http.ListenAndServe(portStr, nil)
}

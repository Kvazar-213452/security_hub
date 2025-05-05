package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/core"
	"head/main_com/download_dep"
	"head/main_com/func_all"
	"head/main_com/module"
	"head/main_com/page/register"
	"head/main_com/page/settings"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/pkg/browser"
)

// core/des/main.go

func startCore(port int) error {
	core.Cmd = exec.Command("./main.exe", strconv.Itoa(port))
	core.Cmd.Stdout = os.Stdout
	core.Cmd.Stderr = os.Stderr
	core.Cmd.Dir = "../"

	if err := core.Cmd.Start(); err != nil {
		return fmt.Errorf("error run core: %v", err)
	}

	go func() {
		err := core.Cmd.Wait()
		if err != nil {
			fmt.Printf("core end error: %v\n", err)
		} else {
			fmt.Println("core good end")
		}
		core.Cleanup()
		os.Exit(0)
	}()

	return nil
}

func main() {
	if _, err := os.Stat("../data"); os.IsNotExist(err) {
		fmt.Println("ins")
		download_dep.Start()
	} else {
		fmt.Println("ok")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		<-sigChan
		core.Cleanup()
		os.Exit(0)
	}()

	defer core.Cleanup()

	var port1 int
	port_config := func_all.PrintPortFromConfig()

	if port_config > 0 {
		port1 = port_config
	} else {
		port1 = func_all.FindFreePort()
	}

	if err := startCore(port1); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	var port int = func_all.FindFreePort()

	time.Sleep(1 * time.Second)

	err := module.RunModules(
		"../data/config_module.json",
		"result.json",
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
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
	http.HandleFunc("/config_global", settings.Post_config_global)
	http.HandleFunc("/visualization", settings.Post_config_change)
	http.HandleFunc("/log_change", settings.Post_log_change)
	http.HandleFunc("/port_change", settings.Post_port_change)
	http.HandleFunc("/shell_change", settings.Post_shell_change)
	http.HandleFunc("/change_lang_settings", settings.Post_change_lang_settings)
	http.HandleFunc("/style_change", settings.Post_style_change)
	http.HandleFunc("/send_email", register.Post_send_email)
	http.HandleFunc("/code_verefic", register.Post_code_verefic)
	http.HandleFunc("/reg_file_unix", register.Post_reg_file_unix)
	http.HandleFunc("/updata_app", settings.Post_updata_app)
	http.HandleFunc("/accses_updata", settings.Post_accses_updata)
	http.HandleFunc("/info_module_nm", settings.Post_info_module_nm)
	http.HandleFunc("/install_module", settings.Post_install_module)
	http.HandleFunc("/uninstall_module", settings.Post_uninstall_module)
	http.HandleFunc("/install_style", main_com.Post_install_style)
	http.HandleFunc("/del_temp", settings.Post_del_temp)
	http.HandleFunc("/get_temp_info", settings.Post_get_temp_info)
	http.HandleFunc("/version_get_server", main_com.Post_version_get_server)
	http.HandleFunc("/version_get", main_com.Post_version_get)
	http.HandleFunc("/get_info_work_server_register", main_com.Post_get_info_work_server_register)
	http.HandleFunc("/get_info_work_server_data_file", main_com.Post_get_info_work_server_data_file)
	http.HandleFunc("/log_out", main_com.Post_log_out)
	http.HandleFunc("/login_acaunt", register.Post_login_acaunt)
	http.HandleFunc("/install_module_app", main_com.Post_install_model_app)
	http.HandleFunc("/uinstall_module_app", main_com.Post_uinstall_model_app)
	http.HandleFunc("/reload_model", main_com.Rost_reload_model)
	http.HandleFunc("/url_site_open", main_com.Rost_open_url)

	portStr1 := ":" + strconv.Itoa(port1)

	server := &http.Server{Addr: portStr}

	go func() {
		fmt.Println("server run on", "http://localhost"+portStr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("error run server:", err)
			os.Exit(1)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	visual := func_all.Visual_сonfig()
	if visual == 0 {
		err := browser.OpenURL("http://localhost" + portStr1)
		if err != nil {
			fmt.Println("error open browser:", err)
		}
	}

	select {}
}

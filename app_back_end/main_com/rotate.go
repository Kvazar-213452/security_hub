package main_com

import (
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"html/template"
	"net/http"
	"os"
)

func Render_index_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /index")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/index.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func Render_settings_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /settings")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/settings.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "settings.html", nil)
}

func Render_system_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /system")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/system.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "system.html", nil)
}

func Get_off_app(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /off_app")

	os.Exit(0)
}

func Render_cleaning(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /cleaning")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/cleaning.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "cleaning.html", nil)
}

func Render_antivirus(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /antivirus")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder+"/templates/html/antivirus/site.html",
		config_main.Frontend_folder+"/templates/html/antivirus/file.html",
		config_main.Frontend_folder+"/templates/html/antivirus/background.html",
		config_main.Frontend_folder+"/templates/antivirus.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "antivirus.html", nil)
}

func Render_encryption(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /encryption")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/encryption.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "encryption.html", nil)
}

func Render_wifi_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /wifi")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/wifi.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "wifi.html", nil)
}

func Render_server_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /server")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/server.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "server.html", nil)
}

func Render_password_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /password")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/password.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "password.html", nil)
}

func Render_register_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /register")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/register.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "register.html", nil)
}

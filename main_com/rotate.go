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

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Render_about_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /about")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/about.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "settings.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "system.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "cleaning.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "antivirus.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "encryption.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "wifi.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "server.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err = tmpl.ExecuteTemplate(w, "password.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Render_file_system_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /file_system")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/file_system.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "file_system.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Render_version_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /version")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/version.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "version.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

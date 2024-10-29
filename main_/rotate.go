package main_

import (
	config_main "head/main_/config"
	"head/main_/func_all"
	"html/template"
	"net/http"
	"os"

	"github.com/pkg/browser"
)

func Render_index_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /index")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/index.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/about.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/settings.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/system.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/cleaning.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
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
		config_main.Frontend_folder+"/templates/html/menu.html",
		config_main.Frontend_folder+"/templates/html/console.html",
		config_main.Frontend_folder+"/templates/encryption.html",
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

func Browser_site_app(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /Browser_site_app")

	url := "https://www.youtube.com/watch?v=pU7N9pVCIl0"
	if err := browser.OpenURL(url); err != nil {
		panic(err)
	}
}

package main_

import (
	"head/main_/func_all"
	"html/template"
	"net/http"
	"os"
)

func Render_index_page(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /index")

	tmpl, err := template.ParseFiles(
		"front_end/templates/html/menu.html",
		"front_end/templates/html/console.html",
		"front_end/templates/index.html",
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
		"front_end/templates/html/menu.html",
		"front_end/templates/html/console.html",
		"front_end/templates/about.html",
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
		"front_end/templates/html/menu.html",
		"front_end/templates/html/console.html",
		"front_end/templates/settings.html",
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
		"front_end/templates/html/menu.html",
		"front_end/templates/html/console.html",
		"front_end/templates/system.html",
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

func Post_off_app(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /off_app")

	os.Exit(0)
}

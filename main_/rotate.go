package main_

import (
	"html/template"
	"net/http"
)

func Render_index_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/html/console.html",
		"templates/index.html",
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
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/html/console.html",
		"templates/about.html",
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
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/html/console.html",
		"templates/settings.html",
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
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/html/console.html",
		"templates/system.html",
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

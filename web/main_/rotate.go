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
		"templates/about.pug",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "about.pug", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Render_settings_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/html/console.html",
		"templates/settings.pug",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "settings.pug", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

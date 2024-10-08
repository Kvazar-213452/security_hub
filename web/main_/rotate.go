package main_

import (
	"html/template"
	"net/http"
)

func Render_index_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/html/menu.html",
		"templates/index.pug",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.pug", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

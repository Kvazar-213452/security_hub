package main_com

import (
	config_main "head/main_com/config"
	"html/template"
	"net/http"
)

// app_back_end/main_com/rotate.go

func Render_password_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/password.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "password.html", nil)
}

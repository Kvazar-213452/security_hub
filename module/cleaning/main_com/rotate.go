package main_com

import (
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"html/template"
	"net/http"
)

// app_back_end/main_com/rotate.go

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

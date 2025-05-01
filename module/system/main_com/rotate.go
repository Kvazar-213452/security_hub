package main_com

import (
	config_main "head/main_com/config"
	"html/template"
	"net/http"
)

// module/system/main_com/rotate.go

func Render_system_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/system.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "system.html", nil)
}

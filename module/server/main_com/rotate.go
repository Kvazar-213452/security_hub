package main_com

import (
	config_main "head/main_com/config"
	"html/template"
	"net/http"
)

// module/server/main_com/rotate.go

func Render_server_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/server.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "server.html", nil)
}

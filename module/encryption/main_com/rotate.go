package main_com

import (
	config_main "head/main_com/config"
	"html/template"
	"net/http"
)

// module/encryption/main_com/rotate.go

func Render_encryption(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/encryption.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "encryption.html", nil)
}

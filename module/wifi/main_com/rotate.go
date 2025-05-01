package main_com

import (
	config_main "head/main_com/config"
	"html/template"
	"net/http"
)

// module/wifi/main_com/rotate.go

func Render_wifi_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder + "/templates/wifi.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "wifi.html", nil)
}

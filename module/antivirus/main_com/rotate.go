package main_com

import (
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"html/template"
	"net/http"
)

// app_back_end/main_com/rotate.go

func Render_antivirus(w http.ResponseWriter, r *http.Request) {
	func_all.AppendToLog("transition to /antivirus")

	tmpl, err := template.ParseFiles(
		config_main.Frontend_folder+"/templates/html/antivirus/site.html",
		config_main.Frontend_folder+"/templates/html/antivirus/file.html",
		config_main.Frontend_folder+"/templates/html/antivirus/background.html",
		config_main.Frontend_folder+"/templates/html/antivirus/resource.html",
		config_main.Frontend_folder+"/templates/antivirus.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "antivirus.html", nil)
}

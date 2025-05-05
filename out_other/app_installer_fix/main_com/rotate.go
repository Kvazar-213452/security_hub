package main_com

import (
	"html/template"
	"net/http"
	"os"
)

// core/des/main_com/rotate.go

func Render_main_page(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("main")
	tmpl, err := tmpl.Parse(Html_code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Get_off_app(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

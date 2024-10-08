package main

import (
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:   "Мій сайт на Go",
		Message: "Ласкаво просимо на мій сайт!",
	}

	tmpl, err := template.ParseFiles("templates/index.pug")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

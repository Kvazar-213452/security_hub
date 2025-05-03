package main_com

import (
	"flash_drive/main_com/func_all"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func StartServer(port int) {
	portStr := ":" + strconv.Itoa(port)

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			os.Exit(0)
		} else {
			http.Error(w, "error", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"web/scan.html",
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "scan.html", nil)
	})

	func_all.WriteServerInfo(portStr)

	if err := http.ListenAndServe(portStr, nil); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

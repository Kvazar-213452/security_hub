package main

import (
	"fmt"
	"head/main_com"
	"log"
	"net/http"
	"os/exec"

	"github.com/pkg/browser"
)

func main() {
	version := main_com.Get_version()
	version_conf := main_com.File_config_get_version()
	version_conf = 1

	if version == version_conf {
		cmd := exec.Command("./app_back_end/head.exe")

		err := cmd.Run()
		if err != nil {
			log.Fatal("Error executing the file:", err)
		}

		fmt.Println("File executed successfully")
	} else {
		port := main_com.FindFreePort()
		http.Handle("/", http.FileServer(http.Dir("./html_tmp_update")))

		url := fmt.Sprintf("http://localhost:%d/index.html", port)
		browser.OpenURL(url)

		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}
}

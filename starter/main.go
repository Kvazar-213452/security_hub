package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:5000/version"

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Помилка запиту:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Помилка читання відповіді:", err)
		return
	}

	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Помилка обробки JSON:", err)
		return
	}

	version := response["version"]
	fmt.Println("Версія:", version)
}

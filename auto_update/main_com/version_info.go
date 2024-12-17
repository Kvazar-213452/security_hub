package main_com

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Version int `json:"version"`
}

func Get_version() int {
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	var response map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	versionStr, ok := response["version"].(string)
	if !ok {
		log.Fatal("Version is not a string")
	}

	version, err := strconv.Atoi(versionStr)
	if err != nil {
		log.Fatal("Error converting version to int:", err)
	}

	return version
}

func File_config_get_version() int {
	file, err := os.Open("main_config.json")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	return config.Version
}

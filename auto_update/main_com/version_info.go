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
	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response map[string]interface{}

	json.Unmarshal(body, &response)

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
	file, _ := os.Open(File_config_phat)
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var config Config
	json.Unmarshal(data, &config)

	return config.Version
}

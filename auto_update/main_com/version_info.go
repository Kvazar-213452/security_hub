package main_com

import (
	"bytes"
	"encoding/json"
	"head/main_com/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// auto_update/main_com/version_info.go

type Config struct {
	Version int `json:"version"`
}

func Get_version() int {
	req, _ := http.NewRequest("POST", config.Url, bytes.NewBuffer([]byte{}))

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

	version, _ := strconv.Atoi(versionStr)

	return version
}

func File_config_get_version() int {
	file, _ := os.Open(config.File_config_phat)
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var config Config
	json.Unmarshal(data, &config)

	return config.Version
}

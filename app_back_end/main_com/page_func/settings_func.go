package page_func

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"io/ioutil"
	"os"
	"strconv"
)

// app_back_end/main_com/page_func/settings_func.go

func LoadConfig() *config_main.Config_global {
	file, _ := os.Open(config_main.Main_config)
	defer file.Close()

	var config config_main.Config_global
	json.NewDecoder(file).Decode(&config)

	return &config
}

func UpdateVisualization(newVisualization string, key string) {
	file, _ := os.Open(config_main.Main_config)

	defer file.Close()

	var config config_main.Config_global

	json.NewDecoder(file).Decode(&config)

	if key == "Visualization" {
		newVis, _ := strconv.Atoi(newVisualization)

		config.Visualization = newVis
	}

	outputFile, _ := os.Create(config_main.Main_config)
	defer outputFile.Close()

	json.NewEncoder(outputFile).Encode(config)
}

func LoadConfig1(filename string) *config_main.Config_global {
	file, _ := os.Open(filename)
	defer file.Close()

	var config config_main.Config_global
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	return &config
}

func SaveConfig(filename string, config *config_main.Config_global) {
	data, _ := json.MarshalIndent(config, "", "  ")

	ioutil.WriteFile(filename, data, 0644)
}

func UpdateConfigKey(key, value string) {
	filename := config_main.Main_config

	config := LoadConfig1(filename)

	switch key {
	case "log":
		config.Log, _ = strconv.Atoi(value)
	case "shell":
		config.Shell, _ = strconv.Atoi(value)
	case "port":
		config.Port, _ = strconv.Atoi(value)
	case "lang":
		config.Lang = value
	case "style":
		config.Style = value
	case "antivirus_flash_drive":
		config.Antivirus.Antivirus_flash_drive, _ = strconv.Atoi(value)
	case "antivirus_flash_drive_cmd":
		config.Antivirus.Antivirus_flash_drive_cmd = value
	default:
		fmt.Errorf("невідомий ключ: %s", key)
	}

	SaveConfig(filename, config)
}

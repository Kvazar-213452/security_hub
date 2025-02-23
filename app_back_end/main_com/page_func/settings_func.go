package page_func

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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
		fmt.Errorf("error key: %s", key)
	}

	SaveConfig(filename, config)
}

func remove_key_array(index int, keyToRemove string) int {
	filePath := "../data/module_config.json"

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0
	}

	var data [][]string
	if err := json.Unmarshal(fileData, &data); err != nil {
		return 0
	}

	if index < 0 || index >= len(data) {
		return 0
	}

	for j, key := range data[index] {
		if key == keyToRemove {
			data[index] = append(data[index][:j], data[index][j+1:]...)
			break
		}
	}

	updatedData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return 0
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		return 0
	}

	return 1
}

func add_key_array(index int, key string) int {
	filePath := "../data/module_config.json"

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0
	}

	var data [][]string
	if err := json.Unmarshal(fileData, &data); err != nil {
		return 0
	}

	if index < 0 || index >= len(data) {
		return 0
	}

	for _, existingKey := range data[index] {
		if existingKey == key {
			return 0
		}
	}

	data[index] = append(data[index], key)

	updatedData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return 0
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		return 0
	}

	return 1
}

func Install_NM(val string) int {
	func_all.ClearDirectory("../data/temp")

	url := "http://localhost:5000/static/" + val + ".zip"

	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	cacheDir := "../data/temp"
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return 0
	}

	zipPath := filepath.Join(cacheDir, val+".zip")
	outFile, err := os.Create(zipPath)
	if err != nil {
		return 0
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return 0
	}

	coreDir := "../core"
	err = os.MkdirAll(coreDir, os.ModePerm)
	if err != nil {
		return 0
	}

	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return 0
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		fpath := filepath.Join(coreDir, file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, file.Mode())
			if err != nil {
				return 0
			}
			continue
		}

		inFile, err := file.Open()
		if err != nil {
			return 0
		}
		defer inFile.Close()

		outFile, err := os.Create(fpath)
		if err != nil {
			return 0
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return 0
		}
	}

	if remove_key_array(0, val) == 0 {
		return 0
	}

	if add_key_array(1, val) == 0 {
		return 0
	}

	return 1
}

func Uninstall_NM(val string) int {
	func_all.ClearDirectory("../core/" + val)

	if func_all.Remove_folder("../core/"+val) == 0 {
		return 0
	}

	if remove_key_array(1, val) == 0 {
		return 0
	}

	if add_key_array(0, val) == 0 {
		return 0
	}

	return 1
}

package page_func

import (
	"encoding/json"
	"fmt"
	config_main "head/main_/config"
	"io/ioutil"
	"os"
	"strconv"
)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	URL           string `json:"url"`
	Port          int    `json:"port"`
	Server        string `json:"server"`
}

func LoadConfig() (*Config_global, error) {
	file, err := os.Open(config_main.Main_config)
	if err != nil {
		return nil, fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("не вдалося декодувати JSON: %w", err)
	}

	return &config, nil
}

func UpdateVisualization(newVisualization string, key string) error {
	file, err := os.Open(config_main.Main_config)
	if err != nil {
		return fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return fmt.Errorf("не вдалося декодувати JSON: %w", err)
	}

	if key == "Visualization" {
		newVis, err := strconv.Atoi(newVisualization)
		if err != nil {
			return fmt.Errorf("не вдалося перетворити visualization: %w", err)
		}
		config.Visualization = newVis
	} else if key == "URL" {
		config.URL = newVisualization
	}

	outputFile, err := os.Create(config_main.Main_config)
	if err != nil {
		return fmt.Errorf("не вдалося створити файл для запису: %w", err)
	}
	defer outputFile.Close()

	if err := json.NewEncoder(outputFile).Encode(config); err != nil {
		return fmt.Errorf("не вдалося закодувати JSON: %w", err)
	}

	return nil
}

func LoadConfig1(filename string) (*Config_global, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("не вдалося розпарсити JSON: %w", err)
	}

	return &config, nil
}

func SaveConfig(filename string, config *Config_global) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("не вдалося серіалізувати JSON: %w", err)
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("не вдалося записати у файл: %w", err)
	}

	return nil
}

func UpdateConfigKey(key, value string) error {
	filename := config_main.Main_config

	config, err := LoadConfig1(filename)
	if err != nil {
		return err
	}

	switch key {
	case "log":
		config.Log, err = strconv.Atoi(value)
	case "port":
		config.Port, err = strconv.Atoi(value)
	case "server":
		config.Server = value
	default:
		return fmt.Errorf("невідомий ключ: %s", key)
	}

	return SaveConfig(filename, config)
}

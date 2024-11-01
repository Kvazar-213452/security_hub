package page_func

import (
	"encoding/json"
	"fmt"
	config_main "head/main_/config"
	"os"
	"strconv"
)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	URL           string `json:"url"`
	Port          int    `json:"port"`
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

package main_

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Lang          string `json:"lang"`
	URL           string `json:"url"`
}

func UpdateVisualization(newVisualization int) error {
	file, err := os.Open("data/main_config.json")
	if err != nil {
		return fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return fmt.Errorf("не вдалося декодувати JSON: %w", err)
	}

	config.Visualization = newVisualization

	outputFile, err := os.Create("data/main_config.json")
	if err != nil {
		return fmt.Errorf("не вдалося створити файл для запису: %w", err)
	}
	defer outputFile.Close()

	if err := json.NewEncoder(outputFile).Encode(config); err != nil {
		return fmt.Errorf("не вдалося закодувати JSON: %w", err)
	}

	return nil
}

type LogContent struct {
	Log string `json:"log"`
}

func LoadLogFile() ([]byte, error) {
	content, err := ioutil.ReadFile("data/main.log")
	if err != nil {
		return nil, fmt.Errorf("не вдалося прочитати файл: %w", err)
	}

	logContent := LogContent{
		Log: string(content),
	}

	jsonData, err := json.Marshal(logContent)
	if err != nil {
		return nil, fmt.Errorf("не вдалося закодувати в JSON: %w", err)
	}

	return jsonData, nil
}

type Message struct {
	Massege string `json:"massege"`
}

func AppendToLog(message string) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s || %s\n", message, currentTime)

	file, err := os.OpenFile("data/main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("не вдалося записати у файл: %w", err)
	}

	return nil
}

func LoadConfig_() (*Config_global, error) {
	file, err := os.Open("data/main_config.json")
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

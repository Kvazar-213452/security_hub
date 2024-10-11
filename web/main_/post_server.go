package main_

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ConfigChange struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type LogsResponse struct {
	Logs string `json:"logs"`
}

func Other_server_post_change_config(port string, key string, value interface{}) {
	data := ConfigChange{
		Key:   key,
		Value: value,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Помилка при перетворенні даних у JSON:", err)
		return
	}

	url := "http://localhost" + port + "/change_config"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Помилка при відправці запиту:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Помилка: отримано статус %s\n", resp.Status)
		return
	}

	fmt.Println("Запит успішно виконано!")
}

func Other_server_post_get_log(w http.ResponseWriter, port string) {
	url := "http://localhost" + port + "/get_logs"

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		http.Error(w, "Помилка при відправці запиту", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Помилка: отримано статус "+resp.Status, resp.StatusCode)
		return
	}

	var logsResponse LogsResponse
	err = json.NewDecoder(resp.Body).Decode(&logsResponse)
	if err != nil {
		http.Error(w, "Помилка при декодуванні відповіді", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logsResponse)
}

type Message struct {
	Massege string `json:"massege"`
}

func Other_server_post_message(port string, message string) error {
	url := "http://localhost" + port + "/log_post_message"

	data := Message{
		Massege: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("помилка при перетворенні даних у JSON: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("помилка при відправці запиту: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("помилка: отримано статус %s", resp.Status)
	}

	return nil
}

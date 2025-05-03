package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func Run_flash() {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Помилка отримання робочої директорії: %v\n", err)
		return
	}

	starterPath := filepath.Join("script", "flash_drive", "starter.md")
	data, err := ioutil.ReadFile(starterPath)
	if err != nil {
		fmt.Printf("Помилка читання %s: %v\n", starterPath, err)
		return
	}

	url := string(data)

	if isServerRunning(url) {
		return
	}

	exePath := filepath.Join(workingDir, "script", "flash_drive", "flash_drive.exe")

	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		return
	}

	cmd := exec.Command(exePath)
	cmd.Dir = filepath.Join(workingDir, "script", "flash_drive")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x08000200,
		HideWindow:    true,
	}

	err = cmd.Start()
	if err != nil {
		return
	}

	err = cmd.Process.Release()
	if err != nil {
		return
	}
}

func isServerRunning(url string) bool {
	url = strings.TrimSpace(url)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url + "scan")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

type Config struct {
	FlashDrive bool `json:"flash_drive"`
}

func Read_json_config() bool {
	filePath := "data/bg_script.json"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return false
	}

	return config.FlashDrive
}

func Contrary_val_flash() error {
	filePath := "data/bg_script.json"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("errro read file: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("error JSON: %v", err)
	}

	config.FlashDrive = !config.FlashDrive

	newData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("error JSON: %v", err)
	}

	err = ioutil.WriteFile(filePath, newData, 0644)
	if err != nil {
		return fmt.Errorf("erro in file: %v", err)
	}

	return nil
}

func Off_flash() error {
	filePath := filepath.Join("script", "flash_drive", "starter.md")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error read file: %v", err)
	}

	url := strings.TrimSpace(string(data))
	if url == "" {
		return fmt.Errorf("URL none")
	}

	requestBody := []byte(`{}`)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(url+"off", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error post: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

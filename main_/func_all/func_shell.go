package func_all

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Lang          string `json:"lang"`
	URL           string `json:"url"`
}

func LoadConfig_start(filename string) (Config_global, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config_global{}, fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config_global{}, fmt.Errorf("не вдалося декодувати JSON: %w", err)
	}

	return config, nil
}

func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func Write_config_core(port string) {
	content := `name = Security Hub
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}</style><iframe src="http://127.0.0.1` + port + `" frameborder="0"></iframe>`

	file, err := os.OpenFile("start_conf.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Помилка запису у файл: %v\n", err)
		return
	}
}

func StartShellWeb() *exec.Cmd {
	cmd := exec.Command("./shell_web.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Не вдалося запустити shell_web.exe: %v\n", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	doneChan := make(chan error, 1)
	go func() {
		doneChan <- cmd.Wait()
	}()

	go func() {
		select {
		case sig := <-sigChan:
			fmt.Printf("Отримано сигнал %v. Завершуємо...\n", sig)
			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Не вдалося завершити shell_web.exe: %v\n", err)
			}
			os.Exit(0)
		case err := <-doneChan:
			if err != nil {
				fmt.Printf("shell_web.exe завершився з помилкою: %v\n", err)
			} else {
				fmt.Println("shell_web.exe завершився.")
			}
			os.Exit(0)
		}
	}()

	return cmd
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

func UpdateVisualization(newVisualization string, key string) error {
	file, err := os.Open("data/main_config.json")
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

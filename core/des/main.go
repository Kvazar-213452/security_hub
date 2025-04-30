package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type Config struct {
	ModuleUninstall []string `json:"module_uinstall"`
}

type Result struct {
	Data map[string]int `json:"module_uinstall"`
}

func main() {
	// Читаємо конфігураційний файл
	configFile, err := os.ReadFile("../data/config_module.json")
	if err != nil {
		fmt.Printf("Помилка читання файлу: %v\n", err)
		return
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Помилка парсингу JSON: %v\n", err)
		return
	}

	// Створюємо мапу для результатів
	result := Result{
		Data: make(map[string]int),
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i, module := range config.ModuleUninstall {
		wg.Add(1)
		go func(module string, index int) {
			defer wg.Done()

			// Формуємо шлях до exe-файлу
			exePath := filepath.Join("..", "..", "module", module, "head.exe")

			// Запускаємо exe-файл
			cmd := exec.Command(exePath)
			err := cmd.Start()
			if err != nil {
				fmt.Printf("Помилка запуску %s: %v\n", exePath, err)
				return
			}

			fmt.Printf("Запущено %s (PID: %d)\n", exePath, cmd.Process.Pid)

			// Зберігаємо результат
			mutex.Lock()
			result.Data[module] = index
			mutex.Unlock()
		}(module, i)
	}

	wg.Wait()

	// Зберігаємо результати у новий JSON-файл
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("Помилка створення JSON: %v\n", err)
		return
	}

	err = os.WriteFile("result.json", output, 0644)
	if err != nil {
		fmt.Printf("Помилка запису у файл: %v\n", err)
		return
	}

	fmt.Println("Операція завершена. Результати збережено у result.json")
}

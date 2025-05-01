package main_com

import (
	"encoding/json"
	"fmt"
	"head/main_com/func_all"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

type Config struct {
	ModuleUninstall []string `json:"module_uinstall"`
}

type Result struct {
	Data map[string]int `json:"module_uinstall"`
}

func RunModules(configPath, resultPath string) error {
	// Get and store the original working directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("помилка отримання поточної директорії: %v", err)
	}
	defer os.Chdir(originalDir) // Ensure we return to original directory

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("помилка читання файлу: %v", err)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("помилка парсингу JSON: %v", err)
	}

	result := Result{
		Data: make(map[string]int),
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var firstError error
	var errorMutex sync.Mutex

	for _, module := range config.ModuleUninstall {
		wg.Add(1)
		go func(module string) {
			defer wg.Done()

			moduleDir := filepath.Join("..", "..", "module", module)
			exePath := filepath.Join(moduleDir, "head.exe")

			// Change to module directory
			err := os.Chdir(moduleDir)
			if err != nil {
				errorMutex.Lock()
				if firstError == nil {
					firstError = fmt.Errorf("помилка зміни директорії для %s: %v", module, err)
				}
				errorMutex.Unlock()
				return
			}

			port := func_all.FindFreePort()

			cmd := exec.Command("./head.exe", strconv.Itoa(port))
			cmd.Dir = moduleDir
			err = cmd.Start()
			if err != nil {
				errorMutex.Lock()
				if firstError == nil {
					firstError = fmt.Errorf("помилка запуску %s: %v", exePath, err)
				}
				errorMutex.Unlock()
				return
			}

			fmt.Printf("Запущено модуль %s (PID: %d) на порту %d\n",
				module, cmd.Process.Pid, port)

			mutex.Lock()
			result.Data[module] = port
			mutex.Unlock()

			os.Chdir(originalDir)
		}(module)
	}

	wg.Wait()

	if firstError != nil {
		return firstError
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("помилка створення JSON: %v", err)
	}

	err = os.WriteFile(resultPath, output, 0644)
	if err != nil {
		return fmt.Errorf("помилка запису у файл: %v", err)
	}

	fmt.Printf("Операція завершена. Результати збережено у %s\n", resultPath)
	return nil
}

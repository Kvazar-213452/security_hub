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

	for _, module := range config.ModuleUninstall {
		wg.Add(1)
		go func(module string) {
			defer wg.Done()

			moduleDir := filepath.Join("..", "..", "module", module)
			exePath := filepath.Join(moduleDir, "head.exe")

			originalDir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Помилка отримання поточної директорії: %v\n", err)
				return
			}

			err = os.Chdir(moduleDir)
			if err != nil {
				fmt.Printf("Помилка зміни директорії для %s: %v\n", module, err)
				return
			}
			defer os.Chdir(originalDir)

			port := func_all.FindFreePort()

			cmd := exec.Command("./head.exe", strconv.Itoa(port))
			cmd.Dir = moduleDir
			err = cmd.Start()
			if err != nil {
				fmt.Printf("Помилка запуску %s: %v\n", exePath, err)
				return
			}

			fmt.Printf("Запущено модуль %s (PID: %d) на порту %d\n",
				module, cmd.Process.Pid, port)

			mutex.Lock()
			result.Data[module] = port
			mutex.Unlock()
		}(module)
	}

	wg.Wait()

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

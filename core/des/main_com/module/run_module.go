package module

import (
	"encoding/json"
	"fmt"
	"head/main_com/func_all"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
)

type Config struct {
	ModuleUninstall []string `json:"module_uinstall"`
}

type ModuleInfo struct {
	PID  int `json:"pid"`
	Port int `json:"port"`
}

type Result struct {
	Data map[string]ModuleInfo `json:"module_uinstall"`
}

var (
	runningModules = make(map[string]ModuleInfo)
	modulesMutex   sync.Mutex
)

func RunModules(configPath, resultPath string) error {
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("помилка отримання поточної директорії: %v", err)
	}
	defer os.Chdir(originalDir)

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
		Data: make(map[string]ModuleInfo),
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
			result.Data[module] = ModuleInfo{
				PID:  cmd.Process.Pid,
				Port: port,
			}

			// Зберігаємо інформацію про запущені модулі
			modulesMutex.Lock()
			runningModules[module] = ModuleInfo{
				PID:  cmd.Process.Pid,
				Port: port,
			}
			modulesMutex.Unlock()

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

// KillAllModules завершує всі запущені процеси модулів
func KillAllModules() {
	modulesMutex.Lock()
	defer modulesMutex.Unlock()

	for module, info := range runningModules {
		process, err := os.FindProcess(info.PID)
		if err != nil {
			fmt.Printf("Помилка пошуку процесу %s (PID: %d): %v\n", module, info.PID, err)
			continue
		}

		// Надсилаємо сигнал SIGTERM для коректного завершення
		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Printf("Помилка завершення процесу %s (PID: %d): %v\n", module, info.PID, err)

			// Якщо SIGTERM не спрацював, використовуємо SIGKILL
			err = process.Kill()
			if err != nil {
				fmt.Printf("Не вдалося примусово завершити процес %s (PID: %d): %v\n", module, info.PID, err)
			}
		} else {
			fmt.Printf("Процес %s (PID: %d) успішно завершено\n", module, info.PID)
		}
	}

	// Очищаємо мапу запущених модулів
	runningModules = make(map[string]ModuleInfo)
}

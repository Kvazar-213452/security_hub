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
		return fmt.Errorf("errro directory: %v", err)
	}
	defer os.Chdir(originalDir)

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error read: %v", err)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("error JSON: %v", err)
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
					firstError = fmt.Errorf("error directory %s: %v", module, err)
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
					firstError = fmt.Errorf("error run %s: %v", exePath, err)
				}
				errorMutex.Unlock()
				return
			}

			fmt.Printf("start module %s (PID: %d) на порту %d\n",
				module, cmd.Process.Pid, port)

			mutex.Lock()
			result.Data[module] = ModuleInfo{
				PID:  cmd.Process.Pid,
				Port: port,
			}

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
		return fmt.Errorf("error make JSON: %v", err)
	}

	err = os.WriteFile(resultPath, output, 0644)
	if err != nil {
		return fmt.Errorf("error in file: %v", err)
	}

	return nil
}

func KillAllModules() {
	modulesMutex.Lock()
	defer modulesMutex.Unlock()

	for module, info := range runningModules {
		process, err := os.FindProcess(info.PID)
		if err != nil {
			fmt.Printf("none module %s (PID: %d): %v\n", module, info.PID, err)
			continue
		}

		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Printf("error end %s (PID: %d): %v\n", module, info.PID, err)

			err = process.Kill()
			if err != nil {
				fmt.Printf("error end %s (PID: %d): %v\n", module, info.PID, err)
			}
		} else {
			fmt.Printf("module %s (PID: %d) end\n", module, info.PID)
		}
	}

	runningModules = make(map[string]ModuleInfo)
}

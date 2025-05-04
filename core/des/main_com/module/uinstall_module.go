package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

type ModuleInfo1 struct {
	PID  int `json:"pid"`
	Port int `json:"port"`
}

type ResultJSON struct {
	ModuleUninstall map[string]ModuleInfo1 `json:"module_uinstall"`
}

func UninstallModule(name string) error {
	data, err := ioutil.ReadFile("./result.json")
	if err != nil {
		return fmt.Errorf("failed to read result.json: %v", err)
	}

	var result ResultJSON
	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("failed to parse result.json: %v", err)
	}

	if moduleInfo, exists := result.ModuleUninstall[name]; exists {
		if moduleInfo.PID > 0 {
			process, err := os.FindProcess(moduleInfo.PID)
			if err == nil {
				err = process.Signal(syscall.SIGTERM)
				if err != nil {
					process.Kill()
				}
			}
		}
	}

	modulePath := filepath.Join("..", "..", "module", name)
	cleanPath := filepath.Clean(modulePath)
	if !strings.HasPrefix(cleanPath, filepath.Join("..", "..", "module", name)) {
		return fmt.Errorf("invalid module path: potential directory traversal attack")
	}

	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		return fmt.Errorf("module directory does not exist")
	}

	if err := os.RemoveAll(modulePath); err != nil {
		return fmt.Errorf("failed to remove module directory: %v", err)
	}

	return nil
}

func MoveModuleToInstall(name string) error {
	filePath := "../data/config_module.json"

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("помилка читання файлу: %v", err)
	}

	var config ModuleConfig
	if err := json.Unmarshal(fileData, &config); err != nil {
		return fmt.Errorf("помилка парсингу JSON: %v", err)
	}

	found := false
	newUninstall := make([]string, 0)
	for _, module := range config.Uninstall {
		if module == name {
			found = true
			continue
		}
		newUninstall = append(newUninstall, module)
	}

	if !found {
		return fmt.Errorf("модуль '%s' не знайдено у module_uinstall", name)
	}

	alreadyInInstall := false
	for _, module := range config.Install {
		if module == name {
			alreadyInInstall = true
			break
		}
	}

	if !alreadyInInstall {
		config.Install = append(config.Install, name)
	}

	config.Uninstall = newUninstall

	updatedData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("помилка перетворення у JSON: %v", err)
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		return fmt.Errorf("помилка запису файлу: %v", err)
	}

	return nil
}

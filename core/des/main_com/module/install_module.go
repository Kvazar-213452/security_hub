package module

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Install_module(name string) error {
	url := fmt.Sprintf("https://github.com/Kvazar-213452/data/raw/refs/heads/main/%s.zip", name)
	targetDir := filepath.Join("..", "..", "module", name)
	tempZipPath := filepath.Join(targetDir, "temp.zip")

	defer func() {
		retryRemove(tempZipPath, 3, 100*time.Millisecond)
	}()

	if err := os.RemoveAll(targetDir); err != nil {
		return fmt.Errorf("failed to remove existing directory: %v", err)
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %v", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download zip file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download zip file: status %d", resp.StatusCode)
	}

	out, err := os.Create(tempZipPath)
	if err != nil {
		return fmt.Errorf("failed to create temp zip file: %v", err)
	}

	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close()
		return fmt.Errorf("failed to save zip file: %v", err)
	}

	if err := out.Close(); err != nil {
		return fmt.Errorf("failed to close temp zip file: %v", err)
	}

	if err := extractZip(tempZipPath, targetDir); err != nil {
		return fmt.Errorf("failed to extract zip file: %v", err)
	}

	if err := retryRemove(tempZipPath, 3, 100*time.Millisecond); err != nil {
		return fmt.Errorf("failed to remove temp zip file after retries: %v", err)
	}

	return nil
}

func retryRemove(path string, attempts int, delay time.Duration) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = os.Remove(path)
		if err == nil {
			return nil
		}
		if os.IsNotExist(err) {
			return nil
		}
		time.Sleep(delay)
	}
	return fmt.Errorf("after %d attempts, last error: %v", attempts, err)
}

func extractZip(zipPath, targetDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	targetDir = filepath.Clean(targetDir) + string(filepath.Separator)

	for _, f := range r.File {
		fpath := filepath.Join(targetDir, f.Name)

		if !strings.HasPrefix(filepath.Clean(fpath), targetDir) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, f.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", fpath, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return fmt.Errorf("failed to create parent directories for %s: %v", fpath, err)
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", fpath, err)
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return fmt.Errorf("failed to open zip entry %s: %v", f.Name, err)
		}

		if _, err = io.Copy(outFile, rc); err != nil {
			outFile.Close()
			rc.Close()
			return fmt.Errorf("failed to extract file %s: %v", fpath, err)
		}

		if err := outFile.Close(); err != nil {
			rc.Close()
			return fmt.Errorf("failed to close extracted file %s: %v", fpath, err)
		}
		if err := rc.Close(); err != nil {
			return fmt.Errorf("failed to close zip entry %s: %v", f.Name, err)
		}
	}

	return nil
}

type ModuleConfig struct {
	Install   []string `json:"module_install"`
	Uninstall []string `json:"module_uinstall"`
}

func MoveModuleToUninstall(name string) error {
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
	newInstall := make([]string, 0)
	for _, module := range config.Install {
		if module == name {
			found = true
			continue
		}
		newInstall = append(newInstall, module)
	}

	if !found {
		return fmt.Errorf("модуль '%s' не знайдено у module_install", name)
	}

	alreadyInUninstall := false
	for _, module := range config.Uninstall {
		if module == name {
			alreadyInUninstall = true
			break
		}
	}

	if !alreadyInUninstall {
		config.Uninstall = append(config.Uninstall, name)
	}

	config.Install = newInstall

	updatedData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("помилка перетворення у JSON: %v", err)
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		return fmt.Errorf("помилка запису файлу: %v", err)
	}

	return nil
}

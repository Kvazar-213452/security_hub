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
)

func Install_module(name string) error {
	tempDir := "../data/temp"

	err := os.RemoveAll(tempDir)
	if err != nil {
		return fmt.Errorf("error del: %v", err)
	}

	err = os.MkdirAll(tempDir, 0755)
	if err != nil {
		return fmt.Errorf("erro create temp: %v", err)
	}

	url := "https://github.com/Kvazar-213452/data/raw/refs/heads/main/" + name + ".zip"
	zipPath := filepath.Join(tempDir, name+".zip")

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error dwn: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error zip zip: %v", err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)

	dest := "../../module/" + name
	os.MkdirAll(dest, 0755)

	zipReader, _ := zip.OpenReader(zipPath)

	defer zipReader.Close()

	for _, file := range zipReader.File {
		fPath := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return fmt.Errorf("error make folder: %v", err)
		}

		dstFile, _ := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		fileInArchive, _ := file.Open()

		_, err = io.Copy(dstFile, fileInArchive)

		dstFile.Close()
		fileInArchive.Close()

		if err != nil {
			return fmt.Errorf("error copying: %v", err)
		}
	}

	os.Remove(zipPath)

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
		return fmt.Errorf("error read: %v", err)
	}

	var config ModuleConfig
	if err := json.Unmarshal(fileData, &config); err != nil {
		return fmt.Errorf("error JSON: %v", err)
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
		return fmt.Errorf("none '%s' in module_install", name)
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
		return fmt.Errorf("error JSON: %v", err)
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		return fmt.Errorf("error file: %v", err)
	}

	return nil
}

package update

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// auto_update/main_com/update/del.go

func DeleteFolders() error {
	dir := "../"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.Name() == "auto_update" || file.Name() == "data" || file.Name() == "main_config.json" {
			continue
		}

		fullPath := filepath.Join(dir, file.Name())

		if file.IsDir() {
			err := os.RemoveAll(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove directory: %v", err)
			}
		} else {
			err := os.Remove(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove file: %v", err)
			}
		}
	}

	dataDir := filepath.Join(dir, "data")
	filesInData, err := ioutil.ReadDir(dataDir)
	if err != nil {
		return fmt.Errorf("failed to read data directory: %v", err)
	}

	for _, file := range filesInData {
		if file.Name() == "main_config.json" {
			continue
		}

		fullPath := filepath.Join(dataDir, file.Name())
		if file.IsDir() {
			err := os.RemoveAll(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove directory in data: %v", err)
			}
		} else {
			err := os.Remove(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove file in data: %v", err)
			}
		}
	}

	return nil
}

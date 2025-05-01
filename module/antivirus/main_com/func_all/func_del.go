package func_all

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// app_back_end/main_com/func_all/func_del.go

func Remove_file(filePath string) int {
	err := os.Remove(filePath)
	if err != nil {
		return 0
	}

	return 1
}

func ClearDirectory(dir string) error {
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		err := os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func Remove_folder(folderPath string) int {
	err := os.RemoveAll(folderPath)
	if err != nil {
		return 0
	}

	return 1
}

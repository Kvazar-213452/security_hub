package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func get_folder_size(path string) (int64, error) {
	var size int64 = 0

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}
	return size, nil
}

func Get_info() string {
	paths := []string{"data", "../app_front_end/static/data", "../data/encryption", "../data/temp"}
	var totalSize int64

	for _, path := range paths {
		size, err := get_folder_size(path)
		if err != nil {
			fmt.Printf("error %s: %v\n", path, err)
			continue
		}
		totalSize += size
	}

	sizeMB := float64(totalSize) / 1024 / 1024

	return strconv.FormatFloat(sizeMB, 'f', 2, 64) + " MB"
}

package system

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// app_back_end/main_com/page_func/system/scan_phat.go

type FolderInfo struct {
	Path string
	Size int64
}

func isValidFileExtension(filePath string, includeExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))

	if len(includeExtensions) > 0 {
		for _, validExt := range includeExtensions {
			if ext == "."+strings.ToLower(validExt) {
				return true
			}
		}
		return false
	}
	return true
}

func getFolderSize(path string, includeExtensions []string) (int64, error) {
	var totalSize int64
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isValidFileExtension(filePath, includeExtensions) {
			totalSize += info.Size()
		}
		return nil
	})
	return totalSize, err
}

func scanDirectory(root string, wg *sync.WaitGroup, ch chan<- FolderInfo, includeExtensions []string) {
	defer wg.Done()

	size, err := getFolderSize(root, includeExtensions)
	if err != nil {
		return
	}

	ch <- FolderInfo{
		Path: root,
		Size: size,
	}
}

func Run_scan_dir(rootDir string, includeExtensions []string) (float64, [][]string) {
	var folderSizes []FolderInfo

	var wg sync.WaitGroup
	ch := make(chan FolderInfo)

	rootSize, _ := getFolderSize(rootDir, includeExtensions)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() || path == rootDir {
			return nil
		}

		wg.Add(1)
		go scanDirectory(path, &wg, ch, includeExtensions)
		return nil
	})

	if err != nil {
		fmt.Println("error", err)
		return 0.0, nil
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for folder := range ch {
		folderSizes = append(folderSizes, folder)
	}

	sort.Slice(folderSizes, func(i, j int) bool {
		return folderSizes[i].Size > folderSizes[j].Size
	})

	unix := [][]string{}

	for i := 0; i < 10 && i < len(folderSizes); i++ {
		folder := folderSizes[i]
		percentage := (float64(folder.Size) / float64(rootSize)) * 100
		unix = append(unix, []string{folder.Path, fmt.Sprintf("%d", folder.Size), fmt.Sprintf("%.2f", percentage)})
	}

	return float64(rootSize), unix
}

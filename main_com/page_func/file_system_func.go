package page_func

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type FolderInfo struct {
	Path string
	Size int64
}

func isValidFileExtension(filePath string, includeExtensions, excludeExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))

	if len(includeExtensions) > 0 {
		for _, validExt := range includeExtensions {
			if ext == "."+strings.ToLower(validExt) {
				return true
			}
		}
		return false
	}

	if len(excludeExtensions) > 0 {
		for _, excludedExt := range excludeExtensions {
			if ext == "."+strings.ToLower(excludedExt) {
				return false
			}
		}
	}

	return true
}

func getFolderSize(path string, includeExtensions, excludeExtensions []string) (int64, error) {
	var totalSize int64
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isValidFileExtension(filePath, includeExtensions, excludeExtensions) {
			totalSize += info.Size()
		}
		return nil
	})
	return totalSize, err
}

func scanDirectory(root string, wg *sync.WaitGroup, ch chan<- FolderInfo, includeExtensions, excludeExtensions []string) {
	defer wg.Done()

	size, err := getFolderSize(root, includeExtensions, excludeExtensions)
	if err != nil {
		return
	}

	ch <- FolderInfo{
		Path: root,
		Size: size,
	}
}

func Run_scan_dir(rootDir string, includeExtensions []string, excludeExtensions []string) (float64, [][]string) {

	// Файли, які ми враховуємо
	// Файли, які ми ігноруємо

	var folderSizes []FolderInfo

	var wg sync.WaitGroup
	ch := make(chan FolderInfo)

	rootSize, _ := getFolderSize(rootDir, includeExtensions, excludeExtensions)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() || path == rootDir {
			return nil
		}

		wg.Add(1)
		go scanDirectory(path, &wg, ch, includeExtensions, excludeExtensions)
		return nil
	})

	if err != nil {
		fmt.Println("Помилка при обході шляху", err)
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
		fmt.Println(unix)
	}

	return float64(rootSize), unix
}

package page_func

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// Структура для зберігання інформації про папку
type FolderInfo struct {
	Path string
	Size int64
}

// Функція для перевірки чи є розширення файлу в списку дозволених або ігнорованих
func isValidFileExtension(filePath string, includeExtensions, excludeExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))

	// Якщо є список дозволених розширень, перевіряємо його
	if len(includeExtensions) > 0 {
		for _, validExt := range includeExtensions {
			if ext == "."+strings.ToLower(validExt) {
				return true
			}
		}
		return false
	}

	// Якщо є список ігнорованих розширень, перевіряємо його
	if len(excludeExtensions) > 0 {
		for _, excludedExt := range excludeExtensions {
			if ext == "."+strings.ToLower(excludedExt) {
				return false
			}
		}
	}

	// Якщо список розширень не задано, враховуємо всі файли
	return true
}

// Функція для отримання розміру директорії
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

// Функція для сканування директорії в окремому потоці
func scanDirectory(root string, wg *sync.WaitGroup, ch chan<- FolderInfo, includeExtensions, excludeExtensions []string) {
	defer wg.Done()

	// Отримуємо розмір директорії
	size, err := getFolderSize(root, includeExtensions, excludeExtensions)
	if err != nil {
		return
	}

	// Надсилаємо інформацію в канал
	ch <- FolderInfo{
		Path: root,
		Size: size,
	}
}

func Start_s() {
	rootDir := "C:\\msys64\\mingw64" // Замініть на вашу папку для сканування

	// Масиви розширень файлів
	includeExtensions := []string{} // Файли, які ми враховуємо
	excludeExtensions := []string{} // Файли, які ми ігноруємо

	var folderSizes []FolderInfo

	var wg sync.WaitGroup
	ch := make(chan FolderInfo)

	// Скануємо директорії в rootDir
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() || path == rootDir {
			return nil
		}

		wg.Add(1)
		go scanDirectory(path, &wg, ch, includeExtensions, excludeExtensions)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path", err)
		return
	}

	// Закриваємо канал після завершення роботи всіх горутин
	go func() {
		wg.Wait()
		close(ch)
	}()

	for folder := range ch {
		folderSizes = append(folderSizes, folder)
	}

	// Сортуємо папки за розміром
	sort.Slice(folderSizes, func(i, j int) bool {
		return folderSizes[i].Size > folderSizes[j].Size
	})

	// Виводимо топ 10 найбільших папок
	fmt.Printf("Топ 10 найбільших папок у %s:\n", rootDir)
	var totalSize int64
	for i := 0; i < 10 && i < len(folderSizes); i++ {
		totalSize += folderSizes[i].Size
	}

	// Виведення топ 10 папок
	for i := 0; i < 10 && i < len(folderSizes); i++ {
		folder := folderSizes[i]
		percentage := (float64(folder.Size) / float64(totalSize)) * 100
		fmt.Printf("%d. %s - %d байт (%.2f%% від топ 10 обсягу)\n", i+1, folder.Path, folder.Size, percentage)
	}

	// Загальний розмір усіх знайдених папок
	var totalAllSize int64
	for _, folder := range folderSizes {
		totalAllSize += folder.Size
	}

	// Виведення загального розміру та відсотка
	fmt.Printf("\nЗагальний розмір всіх папок: %d байт\n", totalAllSize)
	fmt.Printf("Обсяг топ 10 папок становить %.2f%% від загального розміру всіх папок\n", (float64(totalSize)/float64(totalAllSize))*100)
}

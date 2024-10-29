package antivirus

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Scan_file_virus(nameFilePath string, hashesFilePath string) int {
	sha1Hash, err := calculateSHA1(nameFilePath)
	if err != nil {
		fmt.Println("Помилка при обчисленні SHA1:", err)
		return 0
	}

	hashes, err := readHashes(hashesFilePath)
	if err != nil {
		fmt.Println("Помилка при читанні хешів:", err)
		return 0
	}

	if contains(hashes, sha1Hash) {
		return 1
	} else {
		return 0
	}
}

func calculateSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func readHashes(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("не вдалося отримати дані: статус відповіді %d", resp.StatusCode)
	}

	var hashes []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		hash := strings.TrimSpace(scanner.Text())
		if hash != "" {
			hashes = append(hashes, hash)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hashes, nil
}

func contains(hashes []string, hash string) bool {
	for _, h := range hashes {
		if h == hash {
			return true
		}
	}
	return false
}

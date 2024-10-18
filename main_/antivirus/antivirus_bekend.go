package antivirus

import (
	"crypto/sha1"
	"fmt"
	"io"
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

func readHashes(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hashes []string
	var hash string

	for {
		_, err := fmt.Fscanf(file, "%s\n", &hash)
		if err != nil {
			break
		}
		hashes = append(hashes, strings.TrimSpace(hash))
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

package update

import (
	"archive/zip"
	"fmt"
	"head/main_com/config"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadAndExtract() error {
	downloadPath := "../app_back_end.zip"
	extractPath := "../"

	err := downloadFile(config.Server_updatafile, downloadPath)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	err = unzip(downloadPath, extractPath)
	if err != nil {
		return fmt.Errorf("failed to unzip file: %v", err)
	}

	err = os.Remove(downloadPath)
	if err != nil {
		return fmt.Errorf("failed to remove zip file: %v", err)
	}

	return nil
}

func downloadFile(url, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func unzip(zipFile, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		filePath := filepath.Join(destDir, file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		err := extractFile(file, filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractFile(file *zip.File, destPath string) error {
	inFile, err := file.Open()
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return err
	}

	err = os.Chmod(destPath, file.Mode())
	if err != nil {
		return err
	}

	return nil
}

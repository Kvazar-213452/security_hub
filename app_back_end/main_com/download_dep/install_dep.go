package download_dep

import (
	"archive/zip"
	"fmt"
	config_main "head/main_com/config"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// app_back_end/main_com/download_dep/install_dep.go

func Start() {
	url := config_main.Url_dep
	zipFile := "app_back_end.zip"
	unzipDir := "../"

	err := downloadFile(zipFile, url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	err = unzip(zipFile, unzipDir)
	if err != nil {
		fmt.Println("Error unzipping file:", err)
		return
	}

	fmt.Println("File downloaded and unzipped successfully!")

	filePath := "../data/start.md"

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("1")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}

	return nil
}

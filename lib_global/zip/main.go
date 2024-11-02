package main

import (
	"C"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//export Unzip
func Unzip(zipFile *C.char, destDir *C.char) *C.char {
	goZipFile := C.GoString(zipFile)
	goDestDir := C.GoString(destDir)

	err := unzipFile(goZipFile, goDestDir)
	if err != nil {
		return C.CString(fmt.Sprintf("Error: %v", err))
	}
	return C.CString("Unzipping completed successfully!")
}

func unzipFile(zipFile string, destDir string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for file: %w", err)
		}

		dstFile, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer dstFile.Close()

		srcFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %w", err)
		}
		defer srcFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file contents: %w", err)
		}
	}

	return nil
}

func main() {}

package main_com

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	targetDir     = `C:\security_hub`
	zipURL        = "https://github.com/Kvazar-213452/data/raw/refs/heads/main/main_full.zip"
	zipPath       = targetDir + `\main_full.zip`
	exeTarget     = `C:\security_hub\core\des\head.exe`
	exeWorkingDir = `C:\security_hub\core\des`
)

func Post_install(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		os.RemoveAll(targetDir)
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			fmt.Println("error create directory:", err)
			return
		}

		err = downloadFile(zipPath, zipURL)
		if err != nil {
			fmt.Println("Download error:", err)
			return
		}

		err = unzip(zipPath, targetDir)
		if err != nil {
			fmt.Println("Unzip error:", err)
			return
		}

		os.Remove(zipPath)

		err = createShortcut(exeTarget, exeWorkingDir)
		if err != nil {
			fmt.Println("Shortcut creation failed:", err)
			return
		}

		fmt.Println("Done.")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(0)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
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

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func createShortcut(targetPath, workingDir string) error {
	psCommand := fmt.Sprintf(`
$WScriptShell = New-Object -ComObject WScript.Shell
$Shortcut = $WScriptShell.CreateShortcut("$env:USERPROFILE\Desktop\security_hub.lnk")
$Shortcut.TargetPath = "%s"
$Shortcut.WorkingDirectory = "%s"
$Shortcut.Save()
`, targetPath, workingDir)

	cmd := exec.Command("powershell", "-Command", psCommand)
	return cmd.Run()
}

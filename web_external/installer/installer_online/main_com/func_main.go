package main_com

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var shellWebCmd *exec.Cmd

func Decode_Base64_ToFile(base64Data string, outputFilePath string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func StartShellWeb(port string) *exec.Cmd {
	htmlContent := fmt.Sprintf(`<style>iframe{position: fixed;height: 100%%;width: 100%%;top: 0%%;left: 0%%;}</style><iframe src='http://127.0.0.1%s' frameborder='0'></iframe>`, port)

	args := []string{
		Name,
		Window_h,
		Window_w,
		htmlContent,
	}

	shellWebCmd = exec.Command(Core_web, args...)
	shellWebCmd.Stdout = os.Stdout
	shellWebCmd.Stderr = os.Stderr

	if err := shellWebCmd.Start(); err != nil {
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan

		if shellWebCmd != nil {
			shellWebCmd.Process.Kill()
		}

		os.Exit(0)
	}()

	go func() {
		shellWebCmd.Wait()

		DeleteFile("shell_web.exe")
		DeleteFile("webview.dll")

		os.Exit(0)
	}()

	return shellWebCmd
}

func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")

	if err != nil {
		return 0
	}

	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func DWN_app() {
	directory := "C:/TSW_app"

	os.MkdirAll(directory, 0755)

	url := "http://localhost:3000/main.zip"
	outputPath := filepath.Join("C:", "TSW_app", "main.zip")

	os.MkdirAll(filepath.Dir(outputPath), 0755)

	DownloadFile(url, outputPath)

	os.Stat(outputPath)

	zipFile := outputPath
	destDir := "C:/TSW_app/"

	Unzip(zipFile, destDir)

	create_lnk()

	DeleteFile(zipFile)

	if shellWebCmd != nil {
		err := shellWebCmd.Process.Kill()
		if err != nil {
			fmt.Printf("Помилка завершення процесу shell_web.exe: %v\n", err)
		}
	}

	os.Exit(0)
}

func DownloadFile(url string, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не вдалося отримати файл з %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("неочікуваний статус-код: %d", resp.StatusCode)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не вдалося створити файл %s: %v", outputPath, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("не вдалося записати файл %s: %v", outputPath, err)
	}

	return nil
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)

	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)

		dstFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, rc)
		if err != nil {
			return err
		}

		dstFile.Close()
		rc.Close()
	}

	return nil
}

func create_lnk() {
	desktopPath := os.Getenv("USERPROFILE") + "\\Desktop"
	shortcutPath := filepath.Join(desktopPath, "main.lnk")
	targetPath := "C:\\TSW_app\\main.exe"
	workingDirectory := "C:\\TSW_app"

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	wsh, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return
	}

	defer wsh.Release()

	wshDispatch, err := wsh.QueryInterface(ole.IID_IDispatch)

	if err != nil {
		return
	}

	defer wshDispatch.Release()

	shortCut, err := oleutil.CallMethod(wshDispatch, "CreateShortcut", shortcutPath)

	if err != nil {
		return
	}

	defer shortCut.Clear()

	oleutil.PutProperty(shortCut.ToIDispatch(), "TargetPath", targetPath)
	oleutil.PutProperty(shortCut.ToIDispatch(), "WorkingDirectory", workingDirectory)
	oleutil.PutProperty(shortCut.ToIDispatch(), "Description", "Ярлик до main.exe")

	oleutil.CallMethod(shortCut.ToIDispatch(), "Save")

}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

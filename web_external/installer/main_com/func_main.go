package main_com

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"head/main_com/base64_code"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func Decode_Base64_ToFile(base64Data string, outputFilePath string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("не вдалося декодувати base64: %v", err)
	}

	err = ioutil.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("не вдалося записати файл: %v", err)
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

	cmd := exec.Command(Core_web, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Не вдалося запустити shell_web.exe: %v\n", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Отримано сигнал. Завершуємо програму...")
		if err := cmd.Process.Kill(); err != nil {
			fmt.Printf("Не вдалося завершити shell_web.exe: %v\n", err)
		}
		os.Exit(0)
	}()

	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Printf("shell_web.exe завершився з помилкою: %v\n", err)
		} else {
			fmt.Println("shell_web.exe завершився.")
		}
		os.Exit(0)
	}()

	return cmd
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
	dirPath := "C:/TSW_app"

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Printf("Не вдалося створити папку: %v\n", err)
		return
	}

	fmt.Println("Папку створено:", dirPath)

	base64Data := base64_code.Base64_code
	directory := "C:/TSW_app"
	filename := "main.zip"

	os.MkdirAll(directory, 0755)

	Decode_Base64_ToFile_1(base64Data, directory, filename)

	fmt.Println("Файл успішно створено!")

	zipFile := "C:/TSW_app/main.zip"
	destDir := "C:/TSW_app/"

	if err := Unzip(zipFile, destDir); err != nil {
		fmt.Printf("Помилка: %v\n", err)
	} else {
		fmt.Println("Розархівація завершена успішно!")
	}

	create_lnk()
}

func Decode_Base64_ToFile_1(base64Data string, directory string, filename string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("не вдалося декодувати base64: %v", err)
	}

	outputFilePath := filepath.Join(directory, filename)

	err = ioutil.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("не вдалося записати файл: %v", err)
	}

	return nil
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("не вдалося відкрити zip-файл: %v", err)
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return fmt.Errorf("не вдалося створити каталог: %v", err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return fmt.Errorf("не вдалося створити батьківський каталог: %v", err)
		}

		dstFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("не вдалося створити файл: %v", err)
		}

		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("не вдалося відкрити файл в архіві: %v", err)
		}

		_, err = io.Copy(dstFile, rc)
		if err != nil {
			return fmt.Errorf("не вдалося скопіювати вміст: %v", err)
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
		fmt.Printf("Не вдалося створити WScript.Shell: %v\n", err)
		return
	}
	defer wsh.Release()

	wshDispatch, err := wsh.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		fmt.Printf("Не вдалося отримати IDISPATCH: %v\n", err)
		return
	}
	defer wshDispatch.Release()

	shortCut, err := oleutil.CallMethod(wshDispatch, "CreateShortcut", shortcutPath)
	if err != nil {
		fmt.Printf("Не вдалося створити ярлик: %v\n", err)
		return
	}
	defer shortCut.Clear()

	oleutil.PutProperty(shortCut.ToIDispatch(), "TargetPath", targetPath)
	oleutil.PutProperty(shortCut.ToIDispatch(), "WorkingDirectory", workingDirectory)
	oleutil.PutProperty(shortCut.ToIDispatch(), "Description", "Ярлик до main.exe")

	oleutil.CallMethod(shortCut.ToIDispatch(), "Save")

	fmt.Println("Ярлик успішно створено на робочому столі.")
}

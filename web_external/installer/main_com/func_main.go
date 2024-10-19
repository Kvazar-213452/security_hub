package main_com

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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

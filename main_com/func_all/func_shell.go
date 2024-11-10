package func_all

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

type Config_global struct {
	Visualization int    `json:"visualization"`
	Log           int    `json:"log"`
	URL           string `json:"url"`
	Port          int    `json:"port"`
	Shell         int    `json:"shell"`
}

func LoadConfig_start(filename string) (Config_global, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config_global{}, fmt.Errorf("не вдалося відкрити файл: %w", err)
	}
	defer file.Close()

	var config Config_global
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config_global{}, fmt.Errorf("не вдалося декодувати JSON: %w", err)
	}

	return config, nil
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

func StartShellWeb(port int, type_ int) *exec.Cmd {
	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	os.Chdir("core")

	var cmd *exec.Cmd

	if type_ == 0 {
		os.Chdir("NM1")

		htmlContent := fmt.Sprintf(`%s/about`, strconv.Itoa(port))

		args := []string{
			config_main.Name,
			config_main.Window_h,
			config_main.Window_w,
			htmlContent,
		}

		cmd = exec.Command(config_main.Core_web, args...)
	} else if type_ == 1 {
		os.Chdir("NM2")

		htmlContent := fmt.Sprintf(`%d/about`, port)

		var port_ int = FindFreePort()
		portStr := strconv.Itoa(port_)

		args := []string{
			config_main.Window_w,
			config_main.Window_h,
			htmlContent,
			config_main.Name,
			portStr,
		}

		cmd = exec.Command(config_main.Core_web_NM2, args...)
	}

	defer func() {
		os.Chdir(originalDir)
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Не вдалося запустити shell_web.exe: %v\n", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	doneChan := make(chan error, 1)
	go func() {
		doneChan <- cmd.Wait()
	}()

	go func() {
		select {
		case sig := <-sigChan:
			fmt.Printf("Отримано сигнал %v. Завершуємо...\n", sig)
			if err := cmd.Process.Kill(); err != nil {
				fmt.Printf("Не вдалося завершити shell_web.exe: %v\n", err)
			}
			os.Exit(0)
		case err := <-doneChan:
			if err != nil {
				fmt.Printf("shell_web.exe завершився з помилкою: %v\n", err)
			} else {
				fmt.Println("shell_web.exe завершився.")
			}
			os.Exit(0)
		}
	}()

	return cmd
}

func AppendToLog(message string) error {
	config, err := LoadConfig_start(config_main.Main_config)

	if err != nil {
		fmt.Printf("Не вдалося завантажити конфігурацію: %v\n", err)
		return nil
	}

	if config.Log == 1 {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		logEntry := fmt.Sprintf("%s || %s\n", message, currentTime)

		file, err := os.OpenFile(config_main.Log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("не вдалося відкрити файл: %w", err)
		}
		defer file.Close()

		if _, err := file.WriteString(logEntry); err != nil {
			return fmt.Errorf("не вдалося записати у файл: %w", err)
		}

		return nil
	} else {
		return nil
	}
}

func BytePtrToString(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	slice := make([]byte, 0)
	for {
		if *ptr == 0 {
			break
		}
		slice = append(slice, *ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}
	return string(slice)
}

func Get_phat_global() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return ""
	}

	exeDir := filepath.Dir(exePath)
	return exeDir
}

func Clear_file(filePath string) {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		//pass
	}
	defer file.Close()
}

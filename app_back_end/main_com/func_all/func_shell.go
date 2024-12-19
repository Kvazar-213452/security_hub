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

// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣤⣿⣿⣿⣤⡀⢀⣤⣤⠀⢀⣠⣴⣶⣦⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣾⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣻⣿⣿⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⡿⠟⠉⠉⠉⠉⠙⠛⠿⣿⣿⣿⣿⣟⣶⣾⡏⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⡿⣻⡿⠔⠒⠓⠀⠀⠀⠀⠒⠒⢼⣿⣿⣿⢹⣟⣿⡄⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⢁⣠⣀⣤⡀⠀⠀⠀⠄⣀⣀⡟⢿⣿⣾⣿⣿⣷⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠀⢸⣿⣯⠟⣁⣈⡙⠇⠀⠀⠐⠛⢫⠙⠻⣾⣿⣿⣽⡷⣽⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣨⣿⡿⢾⣷⣶⣿⠀⠀⠀⢠⣿⣤⣽⡄⠸⣿⡿⠿⢇⠈⢳⣆
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣗⢻⡇⠈⠛⠛⠁⠀⠀⠀⠈⠛⠻⠟⠁⢠⣿⣷⣴⡈⡇⠀⠁
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢯⣹⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⡗⣱⠃⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡟⣦⠀⠀⠀⠀⠲⠀⠀⠀⠀⠀⣰⣿⢿⣿⣿⠁⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣧⣿⣷⢦⡀⠀⠀⠀⢀⣀⡤⢚⣿⣿⣾⠋⣿⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠁⢿⡿⣿⣷⣌⡓⠒⠊⢉⣀⣴⣿⣿⡟⢿⠀⠈⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠃⣸⣿⣿⣿⠀⣾⣿⣿⣿⣿⡿⣄⠀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡔⢁⣿⣿⡿⣴⣿⣿⣿⣯⡾⠀⠈⢷⡄⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠋⠀⣸⢿⣽⡿⠻⣿⣿⣿⣿⠇⠀⠀⠀⡿⣆⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠇⠀⣠⣿⣸⣿⣤⣴⣿⣿⣿⣿⠀⠀⠀⢰⠀⠘⡆⠀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠋⢀⣴⣿⣿⢿⣿⣿⣿⣿⠀⣀⡏⠀⠀⠀⢸⠀⠀⢸⡀⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡏⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣤⣤⡇⠀⠀⠀⡜⠀⣠⠋⣧⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⠁⢿⣿⣿⣿⣿⣇⢀⣯⣿⣿⣿⣿⠀⠀⠀⠀⡇⡴⡇⠀⢸⠀⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠎⡞⠀⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⢀⠏⢀⠇⠀⢸⡄⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡎⢀⡷⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⡎⠀⢸⠀⠀⠀⡇⠀
// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠀⢸⠑⢀⣾⣧⣿⣿⣿⣿⡀⢨⣿⣿⣿⣿⠃⠀⡜⠀⠀⡸⠀⠀⠀⣿⡄

func LoadConfig_start(filename string) config_main.Config_global {
	file, _ := os.Open(filename)
	defer file.Close()

	var config config_main.Config_global
	json.NewDecoder(file).Decode(&config)

	return config
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

func StartShellWeb(port int, type_ int, version_ int) *exec.Cmd {
	originalDir, _ := os.Getwd()
	os.Chdir("core")

	var cmd *exec.Cmd

	if type_ == 0 {
		os.Chdir("NM1")

		htmlContent := fmt.Sprintf(`%s/main`, strconv.Itoa(port))

		args := []string{
			config_main.Name,
			config_main.Window_h,
			config_main.Window_w,
			htmlContent,
		}

		cmd = exec.Command(config_main.Core_web, args...)
	} else if type_ == 1 {
		os.Chdir("NM2")

		htmlContent := fmt.Sprintf(`%s/main`, strconv.Itoa(port))

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

	cmd.Start()
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

func AppendToLog(message string) {
	config := LoadConfig_start(config_main.Main_config)

	if config.Log == 1 {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		logEntry := fmt.Sprintf("%s || %s\n", message, currentTime)

		file, _ := os.OpenFile(config_main.Log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()

		file.WriteString(logEntry)
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
	exePath, _ := os.Executable()

	exeDir := filepath.Dir(exePath)
	return exeDir
}

func Clear_file(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
}

func Config_port(data string) {
	Clear_file(config_main.Starter_file)

	file, _ := os.OpenFile(config_main.Starter_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	file.WriteString(data)
}

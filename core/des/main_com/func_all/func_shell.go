package func_all

import (
	"encoding/json"
	"fmt"
	config_main "head/main_com/config"
	"net"
	"os"
	"os/exec"
)

// app_back_end/main_com/func_all/func_shell.go

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

func Clear_file(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
}

func Starter(data string) {
	Clear_file(config_main.Starter_file)

	file, _ := os.OpenFile(config_main.Starter_file, +os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	file.WriteString("http://localhost:" + data + "/")
}

func Updata_app() {
	fmt.Println("start head.exe...")

	cmd := exec.Command("../auto_update/head.exe")
	cmd.Dir = "../auto_update"

	err := cmd.Start()
	if err != nil {
		fmt.Println("error head.exe:", err)
		return
	}

	fmt.Println("head.exe end app")

	go func() {
		os.Exit(0)
	}()
}

func RestartScript() error {
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	os.Exit(0)
	return nil
}

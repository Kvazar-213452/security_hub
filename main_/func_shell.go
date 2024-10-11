package main_

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port int `toml:"port"`
}

type Config_main struct {
	Port int `toml:"port_main"`
}

func LoadConfig(configPath string) Config {
	var config Config

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("none: %v", err)
	}
	defer configFile.Close()

	toml.NewDecoder(configFile).Decode(&config)

	return config
}

func LoadConfig_main(configPath string) Config_main {
	var config Config_main

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("none: %v", err)
	}
	defer configFile.Close()

	toml.NewDecoder(configFile).Decode(&config)

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

func Write_config_core(port string) {
	content := `name = Security Hub
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}</style><iframe src="http://127.0.0.1` + port + `" frameborder="0"></iframe>`

	file, err := os.OpenFile("start_conf.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Помилка відкриття файлу: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Помилка запису у файл: %v\n", err)
		return
	}
}

func StartShellWeb() *exec.Cmd {
	cmd := exec.Command("./shell_web.exe")
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

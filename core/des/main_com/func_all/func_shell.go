package func_all

import (
	config_main "head/main_com/config"
	"net"
	"os"
)

// app_back_end/main_com/func_all/func_shell.go

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

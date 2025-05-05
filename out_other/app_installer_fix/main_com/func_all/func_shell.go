package func_all

import (
	"net"
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

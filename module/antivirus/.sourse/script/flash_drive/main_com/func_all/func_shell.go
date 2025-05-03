package func_all

import (
	"fmt"
	"net"
	"os"
)

func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func WriteServerInfo(portStr string) {
	file, err := os.OpenFile("starter.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("http://localhost%s\n", portStr+"/"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

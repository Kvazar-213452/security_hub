package main

import (
	"C" // Це важливо для правильного експорту
	"net"
)

//export FindFreePort
func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func main() {}

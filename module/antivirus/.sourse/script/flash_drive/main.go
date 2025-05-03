package main

import (
	"flash_drive/main_com"
	"flash_drive/main_com/func_all"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var port int = func_all.FindFreePort()

	go main_com.MonitorUSB(port)

	main_com.StartServer(port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
}

package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func main() {
	MakeConnection()
}

func MakeConnection() {
	addr := net.UnixAddr{
		Name: "/tmp/gosearch.sock",
		Net:  "unix",
	}
	listener, err := net.ListenUnix(addr.Net, &addr)
	if err != nil {
		log.Panicln("Unable to open unix socket:", err)
	}
	defer listener.Close()

	cmd := exec.Command("go", "run", "./collect/collect.go", addr.Name)
	if err := cmd.Start(); err != nil {
		log.Panicln("Unable to run collector:", err)
	}

	conn, err := listener.AcceptUnix()
	if err != nil {
		log.Panicln("Unable to accept the unix connection:", err)
	}

	data := make([]byte, 1024)
	if _, _, err = conn.ReadFromUnix(data); err != nil {
		log.Panicln("Error reading from the socket:", err)
	}
	fmt.Println(string(data))
}

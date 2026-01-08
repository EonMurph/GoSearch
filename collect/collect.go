package main

import (
	"log"
	"net"
	"os"
)

func main() {
	SendInformation()
}

func SendInformation() {
	network := "unix"
	if len(os.Args) < 1 {
		log.Panicln(
			"Not enough command line args given. \n",
			"Must provide the socket path as a command line argument.",
		)
	}
	networkAddress := os.Args[1]
	addr, err := net.ResolveUnixAddr(network, networkAddress)
	if err != nil {
		log.Panicln("Error resolving the socket addr:", err)
	}
	conn, err := net.DialUnix(network, nil, addr)
	if err != nil {
		log.Panicln("Error connecting to the unix socket:", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("testing"))
	if err != nil {
		log.Panicln("Error writing data to socket:", err)
	}
}

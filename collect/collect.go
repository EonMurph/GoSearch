package main

import (
	"log"
	"net"
)

func main() {
	SendInformation()
}

func SendInformation() {
	network := "unix"
	addr, err := net.ResolveUnixAddr(network, "/tmp/gosearch.sock")
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

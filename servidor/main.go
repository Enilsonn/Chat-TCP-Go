package main

import (
	"fmt"
	"net"
	"os"

	servidor "github.com/Enilsonn/Chat-TCP-Go.git/servidor/interno"
)

func main() {
	addr := net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 8080,
	}
	listenner, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "erro: %v", err)
		return
	}
	defer listenner.Close()

	go servidor.BroadCaster()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Fprintf(os.Stdout, "erro: %v", err)
			continue
		}
		go servidor.HandleConn(conn)
	}
}

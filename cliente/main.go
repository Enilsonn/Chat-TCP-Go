package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial(
		"tcp",
		"127.0.0.1:8080",
	)
	if err != nil {
		fmt.Println("erro ao estabelecer conexão:", err)
		return
	}
	defer conn.Close()

	mensagensRecebidas := make(chan []byte)
	go serverReader(conn, mensagensRecebidas)
	go clienteWriter(mensagensRecebidas)

	scan := bufio.NewScanner(os.Stdin)
	for {
		for scan.Scan() {
			texto := scan.Text()
			_, err = conn.Write([]byte(texto + "\n"))
			if err != nil {
				fmt.Println("erro ao ler sua mensagem:", err)
			}
		}
		// cliente se desconectou
		if scan.Err() != nil {
			break
		}
	}
	close(mensagensRecebidas)
}

func serverReader(conn net.Conn, ch chan<- []byte) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("erro ao ler do servidor:", err)
			return
		}
		ch <- buffer[:n]
	}
}

func clienteWriter(ch <-chan []byte) {
	for msg := range ch {
		fmt.Println(string(msg))
	}
}

package servidor

import (
	"fmt"
	"io"
	"net"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	// canal próprio dessa cliente
	ch := make(chan []byte)

	// apenas imprime as mensagens concorrentemente
	go clienteWriter(conn, ch)

	// endereço do cliente
	who := conn.RemoteAddr().String()
	ch <- []byte("Você é " + who)

	// avisando aos demais clientes um novo participante
	mensagens <- []byte(who + " entrou na conversa")
	entrando <- ch

	// mensages do cliente
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				//cliente se desconectou
				break
			}
			conn.Write([]byte(fmt.Sprintln("erro:", err)))
		}
		mensagens <- []byte(who + ": " + string(buffer[:n]))
	}

	saindo <- ch
	mensagens <- []byte(who + " saiu da conversa")
}

func clienteWriter(conn net.Conn, ch <-chan []byte) {
	for msg := range ch {
		conn.Write(msg)
	}
}

package servidor

func BroadCaster() {
	clientes := make(map[cliente]bool)

	for {
		select {
		case msg := <-mensagens:
			// como é um chat em grupo, as mensagens são enviadas
			// a todos os clientes na conexão
			for cli := range clientes {
				cli <- msg
			}

		case cli := <-entrando:
			clientes[cli] = true

		case cli := <-saindo:
			delete(clientes, cli)
			close(cli)
		}
	}
}

package servidor

type cliente chan<- []byte // canal de mensagens de saida

var (
	entrando  = make(chan cliente) // clientes que entram no chat
	saindo    = make(chan cliente) // clientes que saem do chat
	mensagens = make(chan []byte)  // mensagens enviadas pelos clientes
)

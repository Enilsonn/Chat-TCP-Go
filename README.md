# Chat TCP em Go

Este é um projeto de chat simples utilizando TCP em Go. O sistema consiste em um **servidor** que gerencia as conexões e múltiplos **clientes** que podem enviar e receber mensagens em tempo real.

---

## Estrutura do projeto

```
├── cliente
│   └── main.go          # Código do cliente
├── servidor
│   ├── interno          # Código interno do servidor (broadcaster, canais, handleConn)
│   └── main.go          # Código principal do servidor
├── go.mod
└── README.md
```

---

## Como executar

### 1. Executar o servidor

No root do projeto, rode:

```bash
go run ./servidor
```

O servidor ficará aguardando conexões na porta **8080**.

---

### 2. Executar o cliente

Em outra(s) janela(s) do terminal, no root do projeto, rode para abrir um cliente:

```bash
go run ./cliente
```

Digite suas mensagens e pressione `Enter` para enviá-las. Para sair do chat, digite:

```
sair
```

---

## Funcionalidades

- Suporte a múltiplos clientes simultaneamente.
- Envia e recebe mensagens em tempo real.
- Notifica quando um cliente entra ou sai da conversa.

---

## Tecnologias

- Linguagem: Go
- Comunicação: TCP

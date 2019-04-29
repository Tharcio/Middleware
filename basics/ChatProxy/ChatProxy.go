package ChatProxy

import (
	"MiddlewareImplementation/basics/Requestor"
	"MiddlewareImplementation/basics/Shared"
)

type Chat struct {
	IP string
	Port int
	OID int
}

func (c Chat) Send(Msg string, User string) string{
	return Requestor.Invoke(c, "Chat", "Send", Shared.chatMsg{Msg, User}) //converter para string
}
func (c Chat) Listen() [] string{
	return Requestor.Invoke(c, "Chat", "Listen", []) //converter para o tipo desejado na saída
}
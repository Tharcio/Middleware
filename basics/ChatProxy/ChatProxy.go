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

func (c Chat) Send(Msg string, User string){
	return Requestor.Invoke(c, "Chat", "Send", Shared.chatMsg{Msg, User}) //ver qual o tipo de saída e converter para esse tipo
}
func (c Chat) Listen() [] string{
	return Requestor.Invoke(c, "Chat", "Listen", []) //converter para o tipo desejado na saída
}
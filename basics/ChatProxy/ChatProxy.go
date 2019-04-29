package ChatProxy

import (
	"MiddlewareImplementation/basics/Requestor"
	"MiddlewareImplementation/basics/Shared"
)

func (c Shared.AOR) Send(Msg string, User string){
	return Requestor.Invoke(c, "Chat", "Send", Shared.chatMsg{Msg, User});
}
func (c Shared.AOR) Listen(){
	return Requestor.Invoke(c, "Chat", "Listen", []);
}
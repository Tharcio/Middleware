package ChatServer

import (
	. "MiddlewarImplementation/basics/Shared"
	. "MiddlewarImplementation/basics/NamingProxy"
	. "MiddlewarImplementation/basics/FTProxy"
	"reflect"
	"strings"
)
type Chat struct {
	Msgs [] string
}

func (c Chat) Send(Msg string, User string){
	if(strings.HasPrefix(Msg,":f ")){
		var naming_server Shared.AOR{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.lookup("File Transfer")
		return ft.Send(Msg[3:(strings.LastIndex(Msg,"|"))],[]byte(Msg[(strings.LastIndex(a.Msg,"|")):]))
	}else if (strings.HasPrefix(Msg,":d ")){
		var naming_server Shared.AOR{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.lookup("File Transfer")
		return ft.Download(Msg[3:])
	}else if (strings.HasPrefix(Msg,":l ")){
		var naming_server Shared.AOR{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.lookup("File Transfer")
		return ft.List()
	} else {
		c.Msgs[len(c.Msgs)] = Msg + ":" + User;
		return
	}
}
func (c Chat) Listen(){
	return c.Msgs
}
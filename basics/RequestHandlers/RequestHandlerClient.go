package main

import (
	"fmt"
	"log"
	"net"
	"io"
)

type handleItf struct{
	data []byte,
	addressPort string,
	oid string,
	fcall string
}

type MsgsUser struct {
	msgs map[string] [] byte
}

func sendHandler(chatmsg Shared.chatMsg, aor Shared.AOR, fCall string){
	msgUser := MsgsUser{[]}
	var chatMsgByte := []byte(chatmsg)
	//msgRequestBytes,_:= json.Marshal(chatmsg)
	h := handleItf{msgRequestBytes, aor.IP ++ ":" +aor.Port ++ "", aor.oid, fCall}
	msgUser.msgs[chatmsg.user] = h.handler()

}

func listenHandler(){
	
}

func (h *handleItf) handler(fCall string) []byte{
	l, err := net.Dial("tcp", h.addressPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()	
	for {
		// send to socket
		_,errW := conn.Write(json.Marshal(h))
		for errW != nil{
			_,errW := conn.Write(json.Marshal(h))
		}

		reply, _ := bufio.NewReader(conn).ReadString('\n')
		return reply
		
	}
}

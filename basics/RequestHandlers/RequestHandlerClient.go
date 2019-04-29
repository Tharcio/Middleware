package main

import (
	"MiddlewareImplementation/basics/Shared"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"io"
)

type handleReq struct{
	data []byte
	addressPort string
}

type MsgsUser struct {
	msgs map[string] [] byte
}

func sendHandler(chatmsg Shared.chatMsg, aor Shared.AOR){
	msgUser := MsgsUser{[]}
	msgRequestBytes,_:= json.Marshal(chatmsg)
	h := handleReq{msgRequestBytes, aor.IP ++ ":" ++aor.Port ++ ""}
	msgUser.msgs[chatmsg.user] = h.handler()

}

func listenHandler(){
	
}

func (h *handleReq) handler() []byte{
	l, err := net.Dial("tcp", h.addressPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()	
	for {
		// send to socket
		_,errW := conn.Write([]byte(text))
		for errW != nil{
			_,errW := conn.Write([]byte(text))
		}

		reply, _ := bufio.NewReader(conn).ReadString('\n')
		return reply
		
	}
}

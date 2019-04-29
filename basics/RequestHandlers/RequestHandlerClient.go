
package RequestHandlers

import (
	"MiddlewareImplementation/basics/Shared"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type handleItf struct{
	data []byte
	addressPort string
	oid string
	fcall string
}

type MsgsUser struct {
	msgs map[string] [] byte
}

func SendHandler(chatmsg Shared.ChatMsg, aor Shared.AOR, fCall string){
	msgUser := MsgsUser{map[string][]byte{}}
	chatMsgByte, err1 := json.Marshal(chatmsg)
	if err1 != nil {
		fmt.Println("erro ao tentar converter a mensagem em json");
	}
	//msgRequestBytes,_:= json.Marshal(chatmsg)
	h := handleItf{chatMsgByte, aor.IP + ":" + string(aor.Port), string(aor.OID), fCall}
	msgUser.msgs[chatmsg.User] = h.handler(fCall) //me parece desnecessário o argumento da função handler

}

func ListenHandler(){
	
}

func (h *handleItf) handler(fCall string) []byte{
	l, err := net.Dial("tcp", h.addressPort)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()	
	for {
		// send to socket
		marshalled_h, err := json.Marshal(h);
		if (err != nil) {
			fmt.Println("erro ao tentar converter handleItf em json");
		}
		_,errW := l.Write(marshalled_h);
		for errW != nil{
			_,errW = l.Write(marshalled_h)
		}

		reply, _ := bufio.NewReader(l).ReadString('\n') //isso esta lendo uma string, mas o método retorna array de byte, fiz o cast, mas o ideal seria ler diretamente o mesmo tipo do retorno
		return []byte(reply)
		
	}
}

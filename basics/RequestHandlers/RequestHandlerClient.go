
package RequestHandlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type CRH struct{
	AddressPort string
}
func (h CRH) Send(msg [] byte) net.Conn {
	l, err := net.Dial("tcp", h.AddressPort)
	if err != nil {
		log.Fatal(err)
	}
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(len(msg)))
	_,errW1 := l.Write(bs)
	for errW1 != nil{
		_,errW1 = l.Write(msg)
	}
	_,errW2 := l.Write(msg);
	for errW2 != nil{
		_,errW2 = l.Write(msg)
	}
	return l
}
func (h CRH) Recieve(conn net.Conn) [] byte{
	len_bytes := make([]byte, 8)
	conn.Read(len_bytes)
	var length int;
	reader := bytes.NewReader(len_bytes)
	err := binary.Read(reader, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	buffer := make([]byte,length)
	conn.Read(buffer);
	return buffer
}

/*
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
*/
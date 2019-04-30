package RequestHandlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type SRH struct {
	IP string
	Port int
}
func (h SRH) Recieve() ([] byte, net.Conn) {
	l, err := net.Listen("tcp",h.IP+string(h.Port))
	if err !=nil {
		fmt.Print("erro ao tentar escutar a porta "+string(h.Port)+", "+string(err.Error()))
	}
	conn, err2 := l.Accept()
	if err2 !=nil {
		fmt.Print("erro ao tentar ao receber conexão do cliente")
	}
	len_bytes := make([]byte, 8)
	conn.Read(len_bytes)
	var length int;
	reader := bytes.NewReader(len_bytes)
	err3 := binary.Read(reader, binary.LittleEndian, &length)
	if err3 != nil {
		fmt.Println("binary.Read failed:", err3)
	}
	buffer := make([]byte,length)
	conn.Read(buffer);
	return buffer, conn;
}
func (h SRH) Send(msg []byte, conn net.Conn) {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(len(msg)))
	_,errW1 := conn.Write(bs)
	for errW1 != nil{
		_,errW1 = conn.Write(msg)
	}
	_,errW2 := conn.Write(msg);
	for errW2 != nil{
		_,errW2 = conn.Write(msg)
	}
}


/*type handleItf struct{
	data []byte
	addressPort string
}
*/
/*
func (h *handleItf) handlerRecieve(){
	for {
		l, err := net.Listen("tcp", h.addressPort)
		if err != nil {
			log.Fatal(err)
		}
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			// Echo all incoming data.
			mensagem, erro3 := bufio.NewReader(conn).ReadString('\n')
			if erro3 != nil {
				fmt.Println(erro3)
				os.Exit(3)
			}
			msgRecived := handleItf{[],"","",""}
			UmsrMsg := json.Unmarshall(mensagem,&h)


			//c.Write(h.data)
			//c.Close()
			io.Copy(c, c)
			c.Write(h.Invoker.invoke())
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}*/
/*
func (h *handleRep) handlerSend(){
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			// Echo all incoming data.

			//c.Write(h.data)
			//c.Close()
			io.Copy(c, c)
			c.takeCommand()
			fmt.Print(c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
	
}


func main() {
	//handleRep{[],0,":8081",}

}*/

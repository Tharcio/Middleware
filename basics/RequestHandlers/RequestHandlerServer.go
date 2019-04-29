package main

//É necessario função invoke em invoker que retorne []byte 

import (
	"fmt"
	"log"
	"net"
	"io"
	"json"
	"../Invoker"
)

/*type handleItf struct{
	data []byte
	addressPort string
}
*/
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
}
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
*/

func main() {
	//handleRep{[],0,":8081",}
	
}

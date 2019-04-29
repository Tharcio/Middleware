package main

//É necessario função invoke em invoker que retorne []byte 

import (
	"fmt"
	"log"
	"net"
	"io"
	"../Invoker"
)

type handleRep struct{
	data []byte
	size int
	addressPort string
}

func (h *handleRep) handlerRecieve(){
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
	l, err := net.Listen("tcp", h.addressPort)
	if err != nil {
		log.Fatal(err)
	}
}

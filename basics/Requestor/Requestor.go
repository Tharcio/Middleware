package Requestor
import (
	"MiddlewareImplementation/basics/Shared"
	"fmt"
	"log"
	"net"
	"io"
	"../RequestHandlers"
)

func Invoke(remote_object Shared.AOR,typeMsg string, action string, args interface{}) interface{}{ //invoke tem que ser um método só
	//interface{} significa qualquer tipo (usado em reflection)
	if typeMsg == "Chat"{
		if action == "Send"{
			RequestHandlers.sendHandler(c,typeMsg)
		}
		if action =="Listen"{
			RequestHandlers.listenHandler()
		}
	}else if typeMsg == "File Transfer" {
		if action == "Send" {

		}
		if action == "Listen" {

		}
	} else {

	}
}

package Requestor
import (
	"MiddlewareImplementation/basics/Shared"
	"fmt"
	"log"
	"net"
	"io"
	"../RequestHandlers"
)

func Invoke(c Shared.AOR,typeMsg string, action string, c Shared.chatMsg) interface{}{ //invoke tem que ser um método só
	//interface{} significa que o método pode retornar qualquer coisa (usado em reflection)
	if typeMsg == "Chat"{
		if action == "Send"{
			RequestHandlers.sendHandler(c)
		}
		if action =="Listen"{
			RequestHandlers.listenHandler()
		}
	}else{
		if typeMsg == "File Transfer"{
			if action == "Send"{
	
			}
			if action =="Listen"{
	
			}
		}
	}
}
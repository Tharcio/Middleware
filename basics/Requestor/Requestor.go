package Requestor
import (
	"MiddlewareImplementation/basics/Shared"
	"fmt"
	"log"
	"net"
	"io"
	"../RequestHandlers"
)

func Invoke(c Shared.AOR,typeMsg string, action string, c Shared.chatMsg){ //invoke tem que ser um método só
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
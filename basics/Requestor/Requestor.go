import (
	"fmt"
	"log"
	"net"
	"io"
	"../RequestHandler"
)

func Invoke(c Shared.AOR,typeMsg string, action string, c Shared.chatMsg){
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
package Requestor
import (
	"../Shared"
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
			RequestHandlers.SendHandler(args.(Shared.ChatMsg),remote_object,action)
		} else {
			RequestHandlers.ListenHandler()
		}
	}else if typeMsg == "File Transfer" {
		if action == "Send" {

		} else if action == "List" {

		} else {

		}
	} else {
		if action == "Bind" {

		} else if action == "Lookup" {

		} else {

		}
	}
}

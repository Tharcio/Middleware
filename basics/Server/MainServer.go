package main

import (
	."../Invoker"
	"../NamingProxy"
	"../Shared"
	"fmt"
)
func main(){
	IP, port := "127.0.0.1:", 8080
	invoker := Invoker{IP: IP, Port: port}
	go invoker.Invoke()
	naming_server := NamingProxy.NS{IP, port, 0}
	fmt.Println(naming_server.Bind("Naming",Shared.AOR{IP,port,0}))
	fmt.Println(naming_server.Bind("File Transfer",Shared.AOR{IP,port,1}))
	fmt.Println(naming_server.Bind("Chat",Shared.AOR{IP,port,2}))
}
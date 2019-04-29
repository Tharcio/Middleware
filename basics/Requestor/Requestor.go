package Requestor
import (
	"MiddlewareImplementation/basics/Shared"
	"encoding/json"
	"fmt"
	. "MiddlewareImplementation/basics/RequestHandlers"
)

func Invoke(remote_object Shared.AOR,typeMsg string, action string, args interface{}) interface{}{ //invoke tem que ser um método só
	//interface{} significa qualquer tipo (usado em reflection)
	var crh = CRH{remote_object.IP+string(remote_object.Port)}
	msg := Shared.RequestFormat{args,remote_object.OID, action}
	msgBytes, err1 := json.Marshal(msg)
	if err1 != nil {
		fmt.Println("erro ao tentar converter a mensagem em json")
	}
	return crh.Send(msgBytes)
}

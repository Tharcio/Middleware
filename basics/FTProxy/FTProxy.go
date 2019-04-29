package FTProxy

import (
	"MiddlewareImplementation/basics/Requestor"
	"MiddlewareImplementation/basics/Shared"
)

func (ft Shared.AOR) Send(Msg string, File [] byte){
	return Requestor.Invoke(ft.Aor, "FileTransfer", "Send", [Msg, File]);
}
func (ft Shared.AOR) Download(nome string){
	return Requestor.Invoke(ft.Aor, "FileTransfer", "Download", [nome]);
}
func (ft Shared.AOR) List(){
	return Requestor.Invoke(ft.Aor, "FileTransfer", "List", []);
}
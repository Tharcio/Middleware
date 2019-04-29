package FTProxy

import (
	"MiddlewareImplementation/basics/Requestor"
)

type FT struct {
	IP string
	Port int
	OID int
}

func (ft FT) Send(Msg string, File [] byte) string {
	return Requestor.Invoke(ft, "FileTransfer", "Send", [Msg, File]) //retorna o tipo da operação
}
func (ft FT) Download(nome string) []byte {
	return Requestor.Invoke(ft, "FileTransfer", "Download", [nome]); //converter para o tipo de retorno
}
func (ft FT) List() [] string {
	return Requestor.Invoke(ft, "FileTransfer", "List", []) //converter para o tipo de retorno
}
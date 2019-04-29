package FTServer

import "reflect"

type FileTransfer struct {
	Files map[string] [] byte
}
func (ft FileTransfer) Send(Msg string, File [] byte){
	ft.Files[Msg]=File //no maximo retorna o status da operação
}
func (ft FileTransfer) Download(nome string) [] byte{
	return ft.Files[nome]
}
func (ft FileTransfer) List() []string{
	keys := reflect.ValueOf(ft.Files).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
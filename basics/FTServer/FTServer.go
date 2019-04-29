package FTServer

type FileTransfer struct {
	Files map[string] [] byte
}
func (ft FileTransfer) Send(Msg string, File [] byte){
	ft.Files[Msg]=File
}
func (ft FileTransfer) Download(nome string){
	return ft.Files[nome]
}
func (ft FileTransfer) List(){
	keys := reflect.ValueOf(ft.Files).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
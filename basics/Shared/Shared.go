package Shared

type AOR struct {
	IP string
	Port int
	OID int
}

type ChatMsg struct{
	msg string
	user string
}

type FTMsg struct{
	file []byte
	file_name string
}
 type BindMessage struct {
 	service_name string
 	AOR AOR
 }
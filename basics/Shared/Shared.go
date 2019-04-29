package Shared

type AOR struct {
	IP string
	Port int
	OID int
}

type chatMsg struct{
	msg string
	user string
}

type FTMsg struct{
	file []byte
	user string
}
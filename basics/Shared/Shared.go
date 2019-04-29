package Shared

type AOR struct {
	IP string,
	port int,
	oid int
}

type chatMsg struct{
	msg string,
	user string
}

type FTMsg struct{
	file byte[],
	user string
}
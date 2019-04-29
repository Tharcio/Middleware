package NamingProxy

import (
	"MiddlewareImplementation/basics/Requestor"
	"MiddlewareImplementation/basics/Shared"
)

type NS struct {
	IP string
	Port int
	OID int
}

func (o NS) Bind(service_name string, aor Shared.AOR) string {
	return Requestor.Invoke(o,  "Naming", "Bind", [service_name, aor]) // retorna o status da operação
}
func (o NS) Lookup(service_name string) Shared.AOR{
	return Requestor.Invoke(o,  "Naming", "Lookup", [service_name]);
}
func (o NS) List() []string {
	return Requestor.Invoke(o, "Naming", "List", []);
}
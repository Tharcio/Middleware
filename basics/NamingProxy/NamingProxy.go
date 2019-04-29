package NamingProxy

import (
	"MiddlewareImplementation/basics/Requestor"
	"MiddlewareImplementation/basics/Shared"
)

func (o Shared.AOR) bind(service_name string, aor Shared.AOR) {
	return Requestor.Invoke(o.Aor, "Naming", "bind", [service_name, aor]);
}
func (o Shared.AOR) lookup(service_name string) Shared.AOR{
	return Requestor.Invoke(o.Aor, "Naming", "lookup", [service_name]);
}
func (o Shared.AOR) list() {
	return Requestor.Invoke(o.Aor, "Naming", "list", []);
}
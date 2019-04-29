package NameService

import (
	"MiddlewareImplementation/basics/Shared"
	"reflect"
)

type Naming struct {
	Lookup_table map[string] Shared.AOR
}
func (o Naming) Bind(service_name string, aor Shared.AOR) {
	o.Lookup_table[service_name] = aor; //no maximo retorna o status da operação
}
func (o Naming) Lookup(service_name string) Shared.AOR{
	return o.Lookup_table[service_name]
}
func (o Naming) List() []string {
	keys := reflect.ValueOf(o.Lookup_table).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
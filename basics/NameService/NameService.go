package NameService

import "MiddlewareImplementation/basics/Shared"

type Naming struct {
	Lookup_table map[string] Shared.AOR
}
func (o Naming) bind(service_name string, aor Shared.AOR) {
	o.Lookup_table[service_name] = aor;
}
func (o Naming) lookup(service_name string) {
	return o.Lookup_table[service_name]
}
func (o Naming) list() {
	keys := reflect.ValueOf(o.Lookup_table).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
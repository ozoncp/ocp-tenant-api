package tenant

import (
	"strconv"
)

type Tenant struct {
	Id   uint64
	Name string
	Type uint8
}

func (t Tenant) String() string {
	return "Id: " + strconv.FormatUint(t.Id, 10) +
		"; Name: " + t.Name +
		"; Type: " + strconv.FormatUint(uint64(t.Type), 10)
}

func New(id uint64, name string, tenantType uint8) Tenant {
	return Tenant{
		Id:   id,
		Name: name,
		Type: tenantType,
	}
}

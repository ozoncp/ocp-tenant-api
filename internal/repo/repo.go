package repo

import "github.com/ozoncp/ocp-tenant-api/internal/tenant"

type Repo interface {
	AddTenants(task []tenant.Tenant) error
	RemoveTenant(taskId uint64) error
	DescribeTenant(taskId uint64) (*tenant.Tenant, error)
	ListTenants(limit, offset uint64) ([]tenant.Tenant, error)
}

package flusher

import (
	"github.com/ozoncp/ocp-tenant-api/internal/repo"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
)

type Flusher interface {
	Flush(tenants []tenant.Tenant) []tenant.Tenant
}

type flusher struct {
	storage    repo.Repo
	maxPkgSize uint
}

func (f *flusher) Flush(tenants []tenant.Tenant) []tenant.Tenant {
	toFlush := tenant.SplitToButch(tenants, f.maxPkgSize)
	for index, pkg := range toFlush {
		err := f.storage.AddTenants(pkg)
		if err != nil {
			// возвращаем не загруженные данные
			return tenants[index*int(f.maxPkgSize):]
		}
	}
	return nil
}

func New(storage repo.Repo, size uint) Flusher {
	return &flusher{
		storage:    storage,
		maxPkgSize: size,
	}
}

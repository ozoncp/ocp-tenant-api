package flusher

import (
	"context"
	"github.com/ozoncp/ocp-tenant-api/internal/repo"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
)

type Flusher interface {
	Flush(ctx context.Context, tenants []tenant.Tenant) []tenant.Tenant
}

type flusher struct {
	storage    repo.Repo
	maxPkgSize uint64
}

func (f *flusher) Flush(ctx context.Context, tenants []tenant.Tenant) []tenant.Tenant {
	toFlush := tenant.SplitToButch(tenants, f.maxPkgSize)
	for index, pkg := range toFlush {
		count, err := f.storage.AddTenants(ctx, pkg)
		if err != nil {
			// возвращаем не загруженные данные
			return tenants[uint64(index)*uint64(f.maxPkgSize)+count:]
		}
	}
	return nil
}

func New(storage repo.Repo, size uint64) Flusher {
	return &flusher{
		storage:    storage,
		maxPkgSize: size,
	}
}

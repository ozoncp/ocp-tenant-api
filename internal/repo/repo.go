package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
)

type Repo interface {
	AddTenant(ctx context.Context, task tenant.Tenant) (uint64, error)
	UpdateTenant(ctx context.Context, tenants *tenant.Tenant) (bool, error)
	ListTenants(ctx context.Context, limit, offset uint64) ([]tenant.Tenant, error)
	DescribeTenant(ctx context.Context, taskId uint64) (*tenant.Tenant, error)
	RemoveTenant(ctx context.Context, taskId uint64) (bool, error)
}

const (
	tableName = "tenants"
)

type tenantRepo struct {
	db        sqlx.DB
	chunkSize uint
}

func New(db sqlx.DB, chunkSize uint) tenantRepo {
	return tenantRepo{
		db:        db,
		chunkSize: chunkSize,
	}
}

func (r *tenantRepo) AddTenant(ctx context.Context, tenant tenant.Tenant) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("name", "type").
		Values(tenant.Name, tenant.Type).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&tenant.Id)

	if err != nil {
		return 0, err
	}

	return tenant.Id, nil
}

func (r *tenantRepo) UpdateTenant(ctx context.Context, toUpdate *tenant.Tenant) (bool, error) {
	query := sq.Update(tableName).
		Set("id", toUpdate.Id).
		Set("name", toUpdate.Name).
		Set("type", toUpdate.Type).
		Where(sq.Eq{"id": toUpdate.Id}).
		RunWith(r.db).PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

func (r *tenantRepo) DescribeTenant(ctx context.Context, id uint64) (*tenant.Tenant, error) {
	query := sq.Select("id", "name", "type").
		From(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).PlaceholderFormat(sq.Dollar)

	var selectedTenant tenant.Tenant
	err := query.QueryRowContext(ctx).Scan(&selectedTenant.Id, &selectedTenant.Name, &selectedTenant.Type)
	if err != nil {
		return nil, err
	}
	return &selectedTenant, nil
}

func (r *tenantRepo) ListTenants(ctx context.Context, limit, offset uint64) ([]tenant.Tenant, error) {
	query := sq.Select("id", "name", "type").
		From(tableName).
		RunWith(r.db).PlaceholderFormat(sq.Dollar).
		Limit(limit).
		Offset(offset)

	var tenants []tenant.Tenant
	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var selectedTenant tenant.Tenant
		err = rows.Scan(&selectedTenant.Id, &selectedTenant.Name, &selectedTenant.Type)
		if err != nil {
			continue
		}
		tenants = append(tenants, selectedTenant)
	}
	return tenants, nil
}

func (r *tenantRepo) RemoveTenant(ctx context.Context, id uint64) (bool, error) {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected <= 0 {
		return false, nil
	}
	return true, nil
}

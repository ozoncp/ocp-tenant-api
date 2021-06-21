package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tenant-api/internal/producer"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
	"log"
	"time"
)

type Repo interface {
	AddTenant(ctx context.Context, tenant tenant.Tenant) (uint64, error)
	AddTenants(ctx context.Context, tenants []tenant.Tenant) (uint64, error)
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
	producer  producer.Producer
}

func New(db sqlx.DB, chunkSize uint, producer producer.Producer) tenantRepo {
	return tenantRepo{
		db:        db,
		chunkSize: chunkSize,
		producer:  producer,
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

	err = r.producer.Send(producer.Create, tenant.Id, time.Now())
	if err != nil {
		log.Printf("failed to send message about updating a note to kafka: %v", err)
	}

	return tenant.Id, nil
}

func (r *tenantRepo) AddTenants(ctx context.Context, tenants []tenant.Tenant) (uint64, error) {
	var numberOfTenantsCreated uint64 = 0
	chunks := tenant.SplitToButch(tenants, 5)
	for _, chunk := range chunks {

		query := sq.Insert(tableName).
			Columns("Name", "Type").
			RunWith(r.db).
			PlaceholderFormat(sq.Dollar)
		for _, chunkTenant := range chunk {
			query = query.Values(chunkTenant.Name, chunkTenant.Type)
		}

		result, err := query.ExecContext(ctx)
		if err != nil {
			return numberOfTenantsCreated, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return numberOfTenantsCreated, err
		}
		numberOfTenantsCreated = numberOfTenantsCreated + uint64(rowsAffected)
	}
	return uint64(numberOfTenantsCreated), nil
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

	err = r.producer.Send(producer.Update, toUpdate.Id, time.Now())
	if err != nil {
		log.Printf("failed to send message about updating a note to kafka: %v", err)
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
	err = r.producer.Send(producer.Remove, id, time.Now())
	if err != nil {
		log.Printf("failed to send message about updating a note to kafka: %v", err)
	}
	return true, nil
}

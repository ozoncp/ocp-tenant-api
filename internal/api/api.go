package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-tenant-api/internal/repo"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
	desc "github.com/ozoncp/ocp-tenant-api/pkg/ocp-tenant-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const (
	errTenantNotFound = "tenant not found"
)

type api struct {
	desc.UnimplementedOcpTenantApiServer
	repo repo.Repo
}

func NewOcpTenantApi(repo repo.Repo) desc.OcpTenantApiServer {
	return &api{
		repo: repo,
	}
}

func (a *api) CreateTenantV1(ctx context.Context, req *desc.CreateTenantV1Request) (*desc.CreateTenantV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateTenantV1")
	defer span.Finish()
	log.Printf("Create tenant ...")

	if err := req.Validate(); err != nil {
		log.Println("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newTenant := tenant.Tenant{
		Id:   0,
		Name: string(req.Name),
		Type: uint8(req.Type),
	}

	newTenantId, err := a.repo.AddTenant(ctx, newTenant)
	if err != nil {
		log.Println("failed to create tenant")
		return nil, status.Error(codes.Internal, err.Error())
	}
	newTenant.Id = newTenantId
	log.Printf("Create tenant success (id: %d)\n", newTenantId)

	return &desc.CreateTenantV1Response{TenantId: newTenantId}, nil
}

func (a *api) MultiCreateTenantV1(ctx context.Context, req *desc.MultiCreateTenantV1Request) (*desc.MultiCreateTenantV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateTenantV1")
	defer span.Finish()
	log.Printf("MultiCreate tenants ...")

	if err := req.Validate(); err != nil {
		log.Println("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var tenants []tenant.Tenant
	for _, val := range req.Create {
		tenant := &tenant.Tenant{
			Id:   0,
			Name: val.Name,
			Type: uint8(val.Type),
		}
		tenants = append(tenants, *tenant)
	}

	numberOfCreated, err := a.repo.AddTenants(ctx, tenants)

	if err != nil {
		log.Printf("MultiCreate tenants failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("MultiCreate tenants success")

	return &desc.MultiCreateTenantV1Response{SuccessCount: numberOfCreated}, nil
}

func (a *api) DescribeTenantV1(
	ctx context.Context,
	req *desc.DescribeTenantV1Request,
) (*desc.DescribeTenantV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DescribeTenantV1")
	span.SetTag("id", req.TenantId)
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Println("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dbTenant, err := a.repo.DescribeTenant(ctx, req.TenantId)
	if err != nil {
		log.Println("failed to get tenant")
		return nil, status.Error(codes.NotFound, errTenantNotFound)
	}
	return &desc.DescribeTenantV1Response{Tenant: &desc.Tenant{
		Id:   dbTenant.Id,
		Name: dbTenant.Name,
		Type: uint32(dbTenant.Type),
	}}, nil
}

func (a *api) RemoveTenantV1(ctx context.Context, request *desc.RemoveTenantV1Request) (*desc.RemoveTenantV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RemoveTenantV1")
	span.SetTag("id", request.TenantId)
	defer span.Finish()
	log.Printf("Remove tenant (id: %d) ...", request.TenantId)

	if err := request.Validate(); err != nil {
		log.Println("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	removed, err := a.repo.RemoveTenant(ctx, request.TenantId)
	if err != nil {
		log.Println("failed to remove tenant")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !removed {
		log.Println("tenant not fount")
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("Remove tenant (id: %d) success", request.TenantId)

	return &desc.RemoveTenantV1Response{Removed: removed}, nil
}

func (a *api) UpdateTenantV1(ctx context.Context, request *desc.UpdateTenantV1Request) (*desc.UpdateTenantV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CallbackFirstV1")
	span.SetTag("id", request.Tenant.Id)
	defer span.Finish()

	log.Printf("Update tenant (id: %d) ...", request.Tenant.Id)

	if err := request.Validate(); err != nil {
		log.Printf("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tenant := &tenant.Tenant{
		Id:   request.Tenant.Id,
		Name: request.Tenant.Name,
		Type: uint8(request.Tenant.Type),
	}

	updates, err := a.repo.UpdateTenant(ctx, tenant)
	if err != nil {
		log.Printf("failed to update tenant")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !updates {
		log.Printf("tenant not found")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	log.Printf("Update tenant (id: %d) success", request.Tenant.Id)

	return &desc.UpdateTenantV1Response{Updated: true}, nil
}

func (a *api) ListTenantsV1(ctx context.Context, request *desc.ListTenantsV1Request) (*desc.ListTenantsV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ListTenantsV1")
	defer span.Finish()
	log.Printf("List tenants ...")

	if err := request.Validate(); err != nil {
		log.Printf("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tenants, err := a.repo.ListTenants(ctx, request.Limit, request.Offset)

	if err != nil {
		log.Printf("failed to get tenants")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(tenants) == 0 {
		log.Printf("tenant not found")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var tenantsProto []*desc.Tenant

	for _, dbTenant := range tenants {
		tenantProto := &desc.Tenant{
			Id:   dbTenant.Id,
			Name: dbTenant.Name,
			Type: uint32(dbTenant.Type),
		}
		tenantsProto = append(tenantsProto, tenantProto)
	}

	log.Printf("List tenants success")

	return &desc.ListTenantsV1Response{Tenants: tenantsProto}, nil
}

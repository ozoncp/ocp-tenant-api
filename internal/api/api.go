package api

import (
	"context"

	desc "github.com/ozoncp/ocp-tenant-api/pkg/ocp-tenant-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errTenantNotFound = "tenant not found"
)

type api struct {
	desc.UnimplementedOcpTenantApiServer
}

func (a *api) DescribeTenantV1(
	ctx context.Context,
	req *desc.DescribeTenantV1Request,
) (*desc.DescribeTenantV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errTenantNotFound)
	return nil, err
}

func NewOcpTenantApi() desc.OcpTenantApiServer {
	return &api{}
}

// Package health provides health  
package health

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

var _ grpc_health_v1.HealthServer = &checker{}

type checker struct {
}

func (c *checker) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (response *grpc_health_v1.HealthCheckResponse, err error) {

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (c *checker) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {

	return nil
}

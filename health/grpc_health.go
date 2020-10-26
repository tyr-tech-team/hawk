package health

import (
	"context"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type grpcHealthServer struct {
	checkFunc map[string]CheckFunc
}

// CheckFunc -
type CheckFunc func() error

// NewGRPCHealthServer -
func NewGRPCHealthServer(server *grpc.Server, serverCheckFunc map[string]CheckFunc) {
	health := new(grpcHealthServer)
	health.checkFunc = serverCheckFunc
	healthpb.RegisterHealthServer(server, health)
}

func (g *grpcHealthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {

	f, ok := g.checkFunc[in.GetService()]
	if ok {

	}

	if err := g.checkFunc(); err != nil {
		return &healthpb.HealthCheckResponse{
			Status: healthpb.HealthCheckResponse_NOT_SERVING,
		}, nil
	}

	return nil, nil
}

func (g *grpcHealthServer) Watch(*healthpb.HealthCheckRequest, healthpb.Health_WatchServer) error {

	return nil
}

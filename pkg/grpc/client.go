package grpc

import (
	"github.com/tyr-tech-team/hawk/status"
	"google.golang.org/grpc"
)

// NewGRPCClient -
func NewGRPCClient(service string) (*grpc.ClientConn, error) {

	if service == "" {
		return nil, status.RemoteHostNotFound.WithDetail("找不到Service 名稱").Err()
	}

	// connect
	conn, err := grpc.Dial(
		service,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, status.ConnectFailed.WithDetail(err.Error()).Err()
	}

	return conn, nil
}

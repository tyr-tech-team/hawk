package grpc

import (
	"context"

	"google.golang.org/grpc"
)

// MiddleHandler - 中間層方法
type MiddleHandler func(ctx context.Context) (context.Context, error)

// UnaryServerInterceptor - grpc 擷取層
func UnaryServerInterceptor(Func MiddleHandler) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		newCtx, err := Func(ctx)
		if err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

// StreamServiceInterceptor - grpc stream 擷取層
func StreamServiceInterceptor(Func MiddleHandler) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx, err := Func(stream.Context())
		if err != nil {
			return err
		}
		wrapped := WrapServerStream(stream)

		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}

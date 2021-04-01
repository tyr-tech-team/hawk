package grpc

import (
	"context"

	"google.golang.org/grpc"
)

// WrappedServerStream - 擴充 Stream 的 Context
type WrappedServerStream struct {
	grpc.ServerStream
	// WrappedContext - 可擴充的 Context
	WrappedContext context.Context
}

// Context  - 複寫原先的 Context
func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

// WrapServerStream  - 建立新的可擴充 Stream
func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}
	return &WrappedServerStream{ServerStream: stream, WrappedContext: stream.Context()}
}

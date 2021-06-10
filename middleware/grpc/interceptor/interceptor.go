package interceptor

import (
	"context"
	"encoding/json"
	"net"
	"strings"

	"go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	grpc_codes "google.golang.org/grpc/codes"
)

const (
	instrumentationName = "go.opentelemetry.io/otel"
)

// UnaryServerInterceptor -
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		entries, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithValues(ctx, entries...)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx), req)

		tr := otel.Tracer(instrumentationName)
		ctx, span := tr.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		)
		defer span.End()

		resp, err := handler(ctx, req)
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpc_codes.OK))
		}

		return resp, err
	}
}

// StreamServerInterceptor -
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()

		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		entries, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithValues(ctx, entries...)

		tr := otel.Tracer(instrumentationName)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx), nil)
		ctx, span := tr.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		)
		defer span.End()

		err := handler(srv, wrapServerStream(ctx, ss))

		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpc_codes.OK))
		}

		return err
	}
}

// Extract -
func Extract(ctx context.Context, metadata *metadata.MD, opts ...Option) ([]attribute.KeyValue, trace.SpanContext) {
	c := newConfig(opts...)
	ctx = c.Propagators.Extract(ctx, &metadataSupplier{
		metadata: metadata,
	})

	attributeSet := baggage.Set(ctx)

	return (&attributeSet).ToSlice(), trace.SpanContextFromContext(ctx)
}

func newConfig(opts ...Option) config {
	cfg := config{
		Propagators:    otel.GetTextMapPropagator(),
		TracerProvider: otel.GetTracerProvider(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	cfg.Tracer = cfg.TracerProvider.Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(contrib.SemVersion()),
	)

	return cfg
}

func spanInfo(fullMethod, peerAddress string, req interface{}) (string, []attribute.KeyValue) {
	attrs := []attribute.KeyValue{semconv.RPCSystemGRPC}
	name, mAttrs := parseFullMethod(fullMethod)
	attrs = append(attrs, mAttrs...)
	attrs = append(attrs, peerAttr(peerAddress)...)

	if req != nil {
		// requestMetadata
		j, _ := json.Marshal(req)
		attrs = append(attrs, attribute.KeyValue{
			Key:   attribute.Key("request"),
			Value: attribute.StringValue(string(j)),
		})
	}

	return name, attrs
}

func parseFullMethod(fullMethod string) (string, []attribute.KeyValue) {
	name := strings.TrimLeft(fullMethod, "/")
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		// Invalid format, does not follow `/package.service/method`.
		return name, []attribute.KeyValue(nil)
	}

	var attrs []attribute.KeyValue
	if service := parts[0]; service != "" {
		attrs = append(attrs, semconv.RPCServiceKey.String(service))
	}
	if method := parts[1]; method != "" {
		attrs = append(attrs, semconv.RPCMethodKey.String(method))
	}
	return name, attrs
}

func peerAttr(addr string) []attribute.KeyValue {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return []attribute.KeyValue(nil)
	}

	if host == "" {
		host = "127.0.0.1"
	}

	return []attribute.KeyValue{
		semconv.NetPeerIPKey.String(host),
		semconv.NetPeerPortKey.String(port),
	}
}

func peerFromCtx(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	return p.Addr.String()
}

func statusCodeAttr(c grpc_codes.Code) attribute.KeyValue {
	return attribute.KeyValue{
		Key:   attribute.Key("statusCode"),
		Value: attribute.Int64Value(int64(c)),
	}
}

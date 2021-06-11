package interceptor

import (
	"bytes"
	"io/ioutil"

	"github.com/kataras/iris/v12"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	instrumentationName = "go.opentelemetry.io/otel"

	// TRACEID -
	TRACEID = "traceID"
)

// HTTP attributes.
var (
	HTTPRemoteAddr = attribute.Key("http.remote")
	HTTPBody       = attribute.Key("http.body")
)

// Opentelemetry -
func Opentelemetry(opts ...Option) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		requestMetadata, _ := metadata.FromOutgoingContext(ctx.Request().Context())
		metadataCopy := requestMetadata.Copy()

		tr := otel.Tracer(instrumentationName)
		newCtx, span := tr.Start(
			ctx.Request().Context(),
			ctx.FullRequestURI(),
			trace.WithSpanKind(trace.SpanKindClient),
		)

		ctx.ResetRequest(ctx.Request().WithContext(newCtx))

		defer span.End()

		span.SetAttributes(spanInfo(ctx)...)

		// inject to metadata
		Inject(ctx.Request().Context(), &metadataCopy)

		//  merge ctx
		mergeCtx := metadata.NewOutgoingContext(ctx.Request().Context(), metadataCopy)

		// reset iris context
		ctx.ResetRequest(ctx.Request().WithContext(mergeCtx))

		ctx.Next()

		err := ctx.Values().Get("error")

		if err != nil {
			s, _ := status.FromError(err.(error))
			span.SetStatus(codes.Error, s.Message())
		}
	}
}

func spanInfo(ctx iris.Context) []attribute.KeyValue {
	attrs := []attribute.KeyValue{}

	attrs = append(attrs, []attribute.KeyValue{
		{
			Key:   semconv.HTTPMethodKey,
			Value: attribute.StringValue(ctx.Request().Method),
		},
		{
			Key:   semconv.HTTPStatusCodeKey,
			Value: attribute.IntValue(ctx.GetStatusCode()),
		},
		{
			Key:   semconv.HTTPUserAgentKey,
			Value: attribute.StringValue(ctx.GetHeader("User-Agent")),
		},
		{
			Key:   semconv.HTTPURLKey,
			Value: attribute.StringValue(ctx.FullRequestURI()),
		},
		{
			Key:   semconv.HTTPRouteKey,
			Value: attribute.StringValue(ctx.Path()),
		},
		{
			Key:   attribute.Key(HTTPRemoteAddr),
			Value: attribute.StringValue(ctx.RemoteAddr()),
		},
	}...)

	switch ctx.Request().Method {
	case iris.MethodGet, iris.MethodDelete:
	default:
		body, err := ctx.GetBody()
		if err != nil {
			return attrs
		}

		attrs = append(attrs, attribute.KeyValue{
			Key:   HTTPBody,
			Value: attribute.StringValue(string(body)),
		})

		ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	return attrs
}

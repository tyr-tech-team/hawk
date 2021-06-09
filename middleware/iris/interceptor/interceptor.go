package interceptor

import (
	"github.com/kataras/iris/v12"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

const (
	instrumentationName = "go.opentelemetry.io/otel"

	TRACEID = "traceID"
)

// HTTP attributes.
var (
	HTTPStatus     = attribute.Key("http.status")
	HTTPMethod     = attribute.Key("http.method")
	HTTPPath       = attribute.Key("http.path")
	HTTPURL        = attribute.Key("http.URL")
	HTTPRemoteAddr = attribute.Key("http.remote")
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

		span.SetAttributes([]attribute.KeyValue{
			{
				Key:   attribute.Key(HTTPURL),
				Value: attribute.StringValue(ctx.FullRequestURI()),
			},
			{
				Key:   attribute.Key(HTTPPath),
				Value: attribute.StringValue(ctx.Path()),
			},
			{
				Key:   attribute.Key(HTTPRemoteAddr),
				Value: attribute.StringValue(ctx.RemoteAddr()),
			},
			{
				Key:   attribute.Key(HTTPMethod),
				Value: attribute.StringValue(ctx.Request().Method),
			},
			{
				Key:   attribute.Key(TRACEID),
				Value: attribute.StringValue(span.SpanContext().TraceID().String()),
			},
			{
				Key:   attribute.Key(HTTPStatus),
				Value: attribute.IntValue(ctx.GetStatusCode()),
			},
		}...)

		// inject to metadata
		Inject(ctx.Request().Context(), &metadataCopy)

		//  merge ctx
		mergeCtx := metadata.NewOutgoingContext(ctx.Request().Context(), metadataCopy)

		// reset iris context
		ctx.ResetRequest(ctx.Request().WithContext(mergeCtx))

		ctx.Next()
	}
}

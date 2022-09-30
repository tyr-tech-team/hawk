// Package interceptor provides interceptor ﳑ
package interceptor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	grpc_codes "google.golang.org/grpc/codes"
)

const (
	instrumentationName = "go.opentelemetry.io/otel"

	// TRACEID -
	TRACEID = "traceId"
)

// HTTP attributes.
var (
	HTTPRemoteAddr = attribute.Key("http.remote")
	HTTPBody       = attribute.Key("http.body")
	options        = new(Options)
)

// Opentelemetry -
func Opentelemetry(opts ...Opt) context.Handler {
	return func(ctx iris.Context) {
		for _, o := range opts {
			o(options)
		}

		requestMetadata, _ := metadata.FromOutgoingContext(ctx.Request().Context())
		metadataCopy := requestMetadata.Copy()

		tr := otel.Tracer(instrumentationName)
		newCtx, span := tr.Start(
			ctx.Request().Context(),
			ctx.FullRequestURI(),
		)

		ctx.ResetRequest(ctx.Request().WithContext(newCtx))

		defer span.End()

		// set TraceID
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key(TRACEID),
			Value: attribute.StringValue(span.SpanContext().TraceID().String()),
		})

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
			span.SetAttributes(statusCodeAttr(s.Code()))
			return
		}

		span.SetAttributes(statusCodeAttr(grpc_codes.OK))
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
		// 取得 body 後需要重新設定
		defer func() {
			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}()

		// set body to attribute
		value := attribute.StringValue(string(body))

		// find path
		if v, ok := options.secrets[ctx.GetCurrentRoute().Path()]; ok {
			if err := json.Unmarshal(body, v); err != nil {
				return attrs
			}
			value = attribute.StringValue(v.Secret())
		}

		attrs = append(attrs, attribute.KeyValue{
			Key:   HTTPBody,
			Value: value,
		})
	}

	return attrs
}

func statusCodeAttr(c grpc_codes.Code) attribute.KeyValue {
	return attribute.KeyValue{
		Key:   attribute.Key("statusCode"),
		Value: attribute.Int64Value(int64(c)),
	}
}

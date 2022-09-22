package iris

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/tyr-tech-team/hawk/middleware/grpc"
	"github.com/tyr-tech-team/hawk/trace"

	tracing "go.opentelemetry.io/otel/trace"
)

// TraceID -
func TraceID() context.Handler {
	return func(ctx iris.Context) {
		r := ctx.Request()
		rid := fmt.Sprintf("%s", trace.NewTraceID())

		// get with tracing
		span := tracing.SpanFromContext(ctx.Request().Context())
		if ok := span.SpanContext().TraceID().IsValid(); ok {
			rid = span.SpanContext().TraceID().String()
		}

		// set TraceID to context
		nctx := trace.SetTraceID(r.Context(), rid)
		// set TraceID to grpc
		nctx = grpc.AppendTraceID(nctx, rid)

		// reset context
		ctx.ResetRequest(r.WithContext(nctx))

		// set header
		ctx.Header("X-Trace-Id", rid)

		ctx.Next()
	}
}

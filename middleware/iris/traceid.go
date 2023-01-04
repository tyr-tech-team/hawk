package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/tyr-tech-team/hawk/middleware/grpc"
	"github.com/tyr-tech-team/hawk/trace"
)

// TraceID -
func TraceID() context.Handler {
	return func(ctx iris.Context) {
		r := ctx.Request()
		rid := trace.NewTraceID()

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

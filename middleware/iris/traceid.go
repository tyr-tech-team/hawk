package iris

import (
	"fmt"

	"github.com/kataras/iris/v12/context"
	"github.com/tyr-tech-team/hawk/middleware/grpc"
	"github.com/tyr-tech-team/hawk/trace"
)

// TraceID -
func TraceID() context.Handler {
	return func(ctx context.Context) {
		r := ctx.Request()
		rid := fmt.Sprintf("web:%s", trace.NewTraceID())

		nctx := trace.SetTraceID(r.Context(), rid)
		nctx = grpc.AppendTraceID(nctx, rid)

		ctx.ResetRequest(r.WithContext(nctx))
		ctx.Header("X-Trace-Id", rid)

		ctx.Next()
	}
}

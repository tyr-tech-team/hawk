// Package gin provides gin ﳑ
package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/tyr-tech-team/hawk/middleware/grpc"
	"github.com/tyr-tech-team/hawk/trace"
)

// TraceID -
func TraceID(c *gin.Context) {

	ctx := c.Request.Context()
	rid := "g:" + trace.NewTraceID()
	nctx := trace.SetTraceID(ctx, rid)

	nctx = grpc.AppendTraceID(nctx, rid)

	c.Request = c.Request.WithContext(nctx)
	c.Header("X-Trace-Id", rid)
	c.Next()
}

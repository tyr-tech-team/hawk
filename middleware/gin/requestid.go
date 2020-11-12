package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/tyr-tech-team/hawk/middleware/grpc"
	"github.com/tyr-tech-team/hawk/trace"
)

// RequestID -
func RequestID(c *gin.Context) {

	ctx := c.Request.Context()
	rid := "g:" + trace.NewRequestID()
	nctx := trace.SetRequestID(ctx, rid)

	nctx = grpc.AppendRequestID(nctx, rid)
	// GrpcRequestID -

	c.Request = c.Request.WithContext(nctx)
	c.Header("X-Request-Id", rid)
	c.Next()
}

package trace

import (
	"context"

	"github.com/bwmarrin/snowflake"
)

// ContextKey -
type ContextKey string

//
const (
	// TraceID -
	TraceID ContextKey = "traceId"
)

var r traceID

func init() {
	r = traceID{}
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic("init requestid generator failed")
	}
	r.node = node
}

type traceID struct {
	node *snowflake.Node
}

// ID -
func (c traceID) GenerateID() string {
	return r.node.Generate().String()
}

// GetTraceID -
func GetTraceID(ctx context.Context) string {
	data := ctx.Value(string(TraceID))
	if data != nil {
		if id, ok := data.(string); ok {
			return id
		}
	}
	return ""
}

// SetTraceID -
func SetTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, string(TraceID), id)
}

// NewTraceID -
func NewTraceID() string {
	return r.GenerateID()
}

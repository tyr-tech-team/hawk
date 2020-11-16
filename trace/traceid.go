package trace

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/tyr-tech-team/hawk/env"
)

var r traceID

func init() {
	r := new(traceID)
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
	data := ctx.Value(env.TraceID)
	if data != nil {
		if id, ok := data.(string); ok {
			return id
		}
	}
	return ""
}

// SetTraceID -
func SetTraceID(ctx context.Context, id ...string) context.Context {
	rid := ""
	if len(id) > 0 && id[0] != "" {
		rid = id[0]
	} else {
		rid = r.GenerateID()
	}
	return context.WithValue(ctx, env.TraceID, rid)
}

// NewTraceID -
func NewTraceID() string {
	return r.GenerateID()
}

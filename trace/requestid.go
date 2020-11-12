package trace

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/tyr-tech-team/hawk/env"
)

var r requestID

func init() {
	r := new(requestID)
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic("init requestid generator failed")
	}

	r.node = node
}

type requestID struct {
	node *snowflake.Node
}

// ID -
func (c requestID) GenerateID() string {
	return r.node.Generate().String()
}

// GetRequestID -
func GetRequestID(ctx context.Context) string {
	data := ctx.Value(env.RequestID)
	if data != nil {
		if id, ok := data.(string); ok {
			return id
		}
	}
	return ""
}

// SetRequestID -
func SetRequestID(ctx context.Context, id ...string) context.Context {
	rid := ""
	if len(id) > 0 && id[0] != "" {
		rid = id[0]
	} else {
		rid = r.GenerateID()
	}
	return context.WithValue(ctx, env.RequestID, rid)
}

// NewRequestID -
func NewRequestID() string {
	return r.GenerateID()
}

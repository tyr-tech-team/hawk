package response

import (
	"context"

	"github.com/tyr-tech-team/hawk/status"
)

// Error -
func Error(ctx context.Context, err error) Response {
	s := status.ConvertStatus(err)
	return withStatus(ctx, nil, s)
}

func withStatus(ctx context.Context, data interface{}, s status.Status) Response {
	r := Response{
		Data:   data,
		Status: newStatus(ctx, s),
	}
	return r
}

// Resp -
func Resp(ctx context.Context, data interface{}) Response {
	return withStatus(ctx, data, status.NoError)
}

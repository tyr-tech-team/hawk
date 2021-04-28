package grpc

import (
	"context"
	"encoding/json"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/tyr-tech-team/hawk/config"
	"google.golang.org/grpc/metadata"
)

var (
	OPERATOR = "operator"
)

// GetOperator -
func GetOperator(ctx context.Context) (context.Context, error) {
	d := new(config.Operator)
	byteStr := ""
	// 先確認是否有metadata的存在
	md, ok := metadata.FromIncomingContext(ctx)

	if ok {
		value := md.Get(OPERATOR)
		if len(value) > 0 && value[0] != "" {
			byteStr = value[0]
			json.Unmarshal([]byte(value[0]), d)
		}
	}

	nctx := context.WithValue(ctx, OPERATOR, d)

	grpc_ctxtags.Extract(ctx).Set(OPERATOR, d)

	return metadata.AppendToOutgoingContext(nctx, OPERATOR, byteStr), nil
}

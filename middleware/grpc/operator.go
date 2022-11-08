package grpc

import (
	"context"
	"encoding/json"
	"time"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/tyr-tech-team/hawk/ctxtags"
	"google.golang.org/grpc/metadata"
)

// GetOperator -
func GetOperator(ctx context.Context) (context.Context, error) {
	d := new(Operator)
	byteStr := ""
	// 先確認是否有metadata的存在
	md, ok := metadata.FromIncomingContext(ctx)

	key := ctxtags.Operator

	if ok {
		value := md.Get(string(key))
		if len(value) > 0 && value[0] != "" {
			byteStr = value[0]
			json.Unmarshal([]byte(value[0]), d)
		}
	}

	nctx := context.WithValue(ctx, string(key), d)

	grpc_ctxtags.Extract(ctx).Set(string(key), d)

	return metadata.AppendToOutgoingContext(nctx, string(key), byteStr), nil
}

// Operator - 操作者資訊
type Operator struct {
	// Name - 操作者姓名
	Name string `json:"name"`
	// Account - 帳號
	Account string `json:"account"`
	// Identifier - 身份類型
	Identifier int32 `json:"identifier"`
	// Time - 操作時間
	Time time.Time `json:"time"`
}

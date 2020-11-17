package grpc

import (
	"context"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/tyr-tech-team/hawk/env"
	"github.com/tyr-tech-team/hawk/trace"
	"google.golang.org/grpc/metadata"
)

// TraceID - 讀取傳送過來的 metadata 中是否擁有TraceID
func TraceID(ctx context.Context) (context.Context, error) {
	tid := ""
	// 先取得是否有 metadata的存在
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 如果有 metadata 的話，檢查是否有 RequestID
		value := md.Get(env.TraceID.String())
		if len(value) > 0 && value[0] != "" {
			// 如果有RequestID 則覆用
			tid = value[0]
		}
	}

	if tid == "" {
		tid = trace.NewTraceID()
	}
	// 將 TraceID 丟入 Conetext 之中
	nctx := trace.SetTraceID(ctx, tid)

	// TODO: 取代
	grpc_ctxtags.Extract(ctx).Set(env.TraceID.String(), tid)

	// 將 TraceID 附加到 OutgoingContext
	return AppendTraceID(nctx, tid), nil
}

// AppendTraceID -
func AppendTraceID(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, env.TraceID.String(), id)
}

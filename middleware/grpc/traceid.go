package grpc

import (
	"context"

	"github.com/tyr-tech-team/hawk/env"
	"github.com/tyr-tech-team/hawk/trace"
	"google.golang.org/grpc/metadata"
)

// TraceID - 讀取傳送過來的 metadata 中是否擁有TraceID
func TraceID(ctx context.Context) context.Context {
	rid := ""
	// 先取得是否有 metadata的存在
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 如果有 metadata 的話，檢查是否有 RequestID
		value := md.Get(env.TraceID.String())
		if len(value) > 0 && value[0] != "" {
			// 如果有RequestID 則覆用
			rid = value[0]
		}
	}

	if rid == "" {
		rid = trace.NewTraceID()
	}
	// 將 TraceID 丟入 Conetext 之中
	nctx := trace.SetTraceID(ctx, rid)
	//
	// 將 TraceID 附加到 OutgoingContext
	return AppendTraceID(nctx, rid)
}

// AppendTraceID -
func AppendTraceID(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, env.TraceID.String(), id)
}

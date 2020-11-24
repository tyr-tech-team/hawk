package response

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tyr-tech-team/hawk/status"
	"github.com/tyr-tech-team/hawk/trace"
)

// Response -
type Response struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

// Status -
type Status struct {
	// 追蹤碼
	TraceID string `json:"traceID"`
	// 新版狀態碼
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	EMessage string   `json:"emessage"`
	Details  []string `json:"details"`
	// RFC3339 時間
	Time string `json:"time"`
	// Timestamp
	Unix int64 `json:"unix"`
}

// JSON -
func (r Response) JSON() []byte {
	data, _ := json.Marshal(r)
	return data
}

//
func newStatus(ctx context.Context, s status.Status) Status {
	t := time.Now().In(time.Local)
	tid := trace.GetTraceID(ctx)
	ss := Status{
		TraceID: tid,
		Time:    t.Format(time.RFC3339),
		Unix:    t.Unix(),
		Details: make([]string, 0),
	}

	if s != nil {
		ss.Code = s.Code()
		ss.Message = s.Message()
		ss.EMessage = s.EMessage()
		ss.Details = s.Detail()
	}

	return ss
}

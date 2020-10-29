package status

import (
	gs "google.golang.org/grpc/status"
)

type status struct {
	gst  *gs.Status
	body *body
}

// NewStatus - 新增狀態碼
func NewStatus() Status {

	return nil
}

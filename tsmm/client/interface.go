package client

import "context"

type Client interface {
	// AddSchedule - 批次新增排程
	AddSchedule(ctx context.Context, req *AddScheduleReq)
	// UpdateSchedule - 批次更新排程
	UpdateSchedule(ctx context.Context, req *UpdateScheduleReq)
	// CancelSchedule - 批次取消排程
	CancelSchedule(ctx context.Context, req *CancelScheduleReq)
}

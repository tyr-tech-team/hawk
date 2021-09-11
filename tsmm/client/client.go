package client

import (
	"context"
	"encoding/json"

	"github.com/tyr-tech-team/hawk/broker"
)

// AddSchedule - 批次新增排程
func (c *client) AddSchedule(ctx context.Context, req *AddScheduleReq) {
	reqBytes, _ := json.Marshal(req.ScheduleList)
	c.broker.Publish(ADD_SCHEDULE_TOPIC, &broker.Message{
		Header: req.Header,
		Body:   reqBytes,
	})
}

// UpdateSchedule - 批次更新排程
func (c *client) UpdateSchedule(ctx context.Context, req *UpdateScheduleReq) {
	reqBytes, _ := json.Marshal(req.ScheduleList)
	c.broker.Publish(UPDATE_SCHEDULE_TOPIC, &broker.Message{
		Header: req.Header,
		Body:   reqBytes,
	})
}

// CancelSchedule - 批次取消排程
func (c *client) CancelSchedule(ctx context.Context, req *CancelScheduleReq) {
	reqBytes, _ := json.Marshal(req.ScheduleIDList)
	c.broker.Publish(CANCEL_SCHEDULE_TOPIC, &broker.Message{
		Header: req.Header,
		Body:   reqBytes,
	})
}

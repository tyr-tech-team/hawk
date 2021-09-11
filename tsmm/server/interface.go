package server

type Server interface {
	// StopSubscribe - 停止訂閱
	StopSubscribe()
	// AddTopic - 新增監聽Topic
	AddTopic(topic string, fn func(map[string]interface{}, []byte) error)
	// AddScheduleReply - 新增排程
	AddScheduleReply(func(map[string]interface{}, *AddScheduleRes) error)
	// UpdateScheduleReply - 更新排程
	UpdateScheduleReply(func(map[string]interface{}, *UpdateScheduleRes) error)
	// CancelScheduleReply - 取消排程
	CancelScheduleReply(func(map[string]interface{}, *CancelScheduleRes) error)
}

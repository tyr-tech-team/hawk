package server

type Server interface {
	// StopSubscribe - 停止訂閱
	StopSubscribe()
	// AddTopic - 新增自訂監聽通道
	AddTopic(topic string, fn func(map[string]interface{}, []byte) error)
	// AddScheduleReply - 監聽新增排程回覆通道
	AddScheduleReply(func(map[string]interface{}, *AddScheduleRes) error)
	// UpdateScheduleReply - 監聽更新排程回覆通道
	UpdateScheduleReply(func(map[string]interface{}, *UpdateScheduleRes) error)
	// CancelScheduleReply - 監聽取消排程回覆通道
	CancelScheduleReply(func(map[string]interface{}, *CancelScheduleRes) error)
}

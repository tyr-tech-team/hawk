package client

import "time"

// 批次新增排程請求
type AddScheduleReq struct {
	Header       map[string]interface{}
	ScheduleList []*AddScheduleDetail
}

// 批次更新排程請求
type UpdateScheduleReq struct {
	Header       map[string]interface{}
	ScheduleList []*UpdateScheduleDetail
}

// 批次取消排程請求
type CancelScheduleReq struct {
	Header         map[string]interface{}
	ScheduleIDList []string
}

// 新增排程請求
type AddScheduleDetail struct {
	// Nats回傳通道
	Topic string `json:"topic"`
	// HTTP回傳接口 EX:https://192.168.1.111:2020/webhook/schedule
	Webhook string `json:"webhook"`
	// cron包的執行週期設定
	Time string `json:"time"`
	// 啟用時間
	StartTime time.Time `json:"startTime"`
	// 停用時間
	StopTime time.Time `json:"stopTime"`
	// TotalTimes - 總運行次數
	TotalTimes int64 `json:"totalTimes"`
	// TimesType - 是否次數限制
	TimesType int32 `json:"timesType"`
	// 供client端使用辨識行為
	Action string `json:"action"`
	// 排程名稱
	Name string `json:"name"`
	// client端資料
	Data []byte `json:"data"`
	// 備註
	Memo string `json:"memo"`
}

// 更新排程請求
type UpdateScheduleDetail struct {
	// 排程編號
	ScheduleID string `json:"scheduleId"`
	// Nats回傳通道
	Topic string `json:"topic"`
	// HTTP回傳接口 EX:https://192.168.1.111:2020/webhook/schedule
	Webhook string `json:"webhook"`
	// cron包的執行週期設定
	Time string `json:"time"`
	// 啟用時間
	StartTime time.Time `json:"startTime"`
	// 停用時間
	StopTime time.Time `json:"stopTime"`
	// TotalTimes - 總運行次數
	TotalTimes int64 `json:"totalTimes"`
	// TimesType - 是否次數限制
	TimesType int32 `json:"timesType"`
	// 供client端使用辨識行為
	Action string `json:"action"`
	// 排程名稱
	Name string `json:"name"`
	// client端資料
	Data []byte `json:"data"`
	// 備註
	Memo string `json:"memo"`
}

const (
	ADD_SCHEDULE_TOPIC    = "sm-add-schedule"
	UPDATE_SCHEDULE_TOPIC = "sm-update-schedule"
	CANCEL_SCHEDULE_TOPIC = "sm-cancel-schedule"
)

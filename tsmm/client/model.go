package client

import "time"

// 批次新增排程請求
type AddScheduleReq struct {
	// Header - 標頭，供Client端自行放值，於Reply時會回覆相同Header
	Header map[string]interface{}
	// ScheduleList - 排程List
	ScheduleList []*AddScheduleDetail
}

// 批次更新排程請求
type UpdateScheduleReq struct {
	// Header - 標頭，供Client端自行放值，於Reply時會回覆相同Header
	Header map[string]interface{}
	// ScheduleList - 排程List
	ScheduleList []*UpdateScheduleDetail
}

// 批次取消排程請求
type CancelScheduleReq struct {
	// Header - 標頭，供Client端自行放值，於Reply時會回覆相同Header
	Header map[string]interface{}
	// ScheduleIDList - 排程IDList
	ScheduleIDList []string
}

// 新增排程請求
type AddScheduleDetail struct {
	// Nats回傳通道
	Topic string `json:"topic"`
	// HTTP回傳接口 EX:https://192.168.1.111:2020/webhook/schedule
	Webhook string `json:"webhook"`
	// cron包的執行週期設定 詳見https://pkg.go.dev/github.com/robfig/cron/v3#hdr-CRON_Expression_Format
	Time string `json:"time"`
	// 啟用時間（需小於第一次運行）
	StartTime time.Time `json:"startTime"`
	// 停用時間（需大於於最後一次運行）
	StopTime time.Time `json:"stopTime"`
	// TotalTimes - 總運行次數
	TotalTimes int64 `json:"totalTimes"`
	// TimesType - 是否次數限制 1: 限制 2:不限制
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
	// cron包的執行週期設定 詳見https://pkg.go.dev/github.com/robfig/cron/v3#hdr-CRON_Expression_Format
	Time string `json:"time"`
	// 啟用時間（需小於第一次運行）
	StartTime time.Time `json:"startTime"`
	// 停用時間（需大於於最後一次運行）
	StopTime time.Time `json:"stopTime"`
	// TotalTimes - 總運行次數
	TotalTimes int64 `json:"totalTimes"`
	// TimesType - 是否次數限制 1: 限制 2:不限制
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

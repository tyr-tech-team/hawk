package event

import "time"

type EventStore interface {
	// 取得事件編號(traceID)
	GetEventID() string

	// 取得事件類型(ex:Create、Update)
	GetEventType() int32

	// 取得事件對象
	GetAggregateID() string

	// 取得對象類別
	GetAggregateType() string

	// 取得事件資料
	GetEventData() string

	// 取得建立時間
	GetCreatedAt() time.Time

	// 取得版本
	GetVersion() string
}

type EventType int32

const (
	// 新增
	EVENT_CREATE EventType = iota + 1
	// 更新
	EVENT_UPDATE
	// 刪除
	EVENT_DELETE
)

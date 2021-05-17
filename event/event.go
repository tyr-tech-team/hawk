package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tyr-tech-team/hawk/trace"
)

func SetEventInterface(ctx context.Context, in EventStore, AggregateType EventType) []byte {
	tid := trace.GetTraceID(ctx)

	e := &Event{
		EventID:       tid,
		EventType:     AggregateType,
		AggregateID:   in.GetAggregateID(),
		AggregateType: in.GetAggregateType(),
		EventData:     in.GetEventData(),
		CreatedAt:     in.GetCreatedAt(),
	}

	j, _ := json.Marshal(e)

	return j
}

func SetEventStruct(ctx context.Context, aggregateId, aggregateType string, eventData interface{}, eventType EventType) ([]byte, error) {

	tid := trace.GetTraceID(ctx)
	d, _ := json.Marshal(eventData)

	e := &Event{
		EventID:       tid,
		EventType:     EventType(eventType),
		AggregateID:   aggregateId,
		AggregateType: aggregateType,
		EventData:     string(d),
		CreatedAt:     time.Now().In(time.Local),
	}

	j, _ := json.Marshal(e)

	return j, nil
}

// EventStore -
type Event struct {
	// 事件編號(traceID)
	EventID string `bson:"eventId"`

	// 事件類型(CRUD)
	EventType EventType `bson:"eventType"`

	// 事件對象
	AggregateID string `bson:"aggregateId"`

	// 對象類別
	AggregateType string `bson:"aggregateType"`

	// 事件資料
	EventData string `bson:"eventData"`

	// 建立時間
	CreatedAt time.Time `bson:"createdAt"`
}

package main

import (
	"encoding/json"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/tyr-tech-team/hawk/broker"
	"github.com/tyr-tech-team/hawk/broker/natsstreaming"
)

func main() {
	n := natsstreaming.New(
		natsstreaming.SetURL(natsstreaming.DefaultURL),
		natsstreaming.SetStanClusterID("test"),
	)

	b := broker.NewBroker(n)

	b.Subscribe("a", func(b broker.Event) error {
		//m := b.Message()
		spew.Dump("Get")

		return nil

	}, broker.Queue("member"))

	time.Sleep(20 * time.Second)

	e := &EventStore{
		EventID:       "1",
		EventType:     2,
		AggregateID:   "278363",
		AggregateType: "memebr",
		EventData:     "7126312",
		CreatedAt:     time.Now().In(time.Local),
	}
	j, _ := json.Marshal(e)
	b.Publish("a", &broker.Message{
		//Body:   []byte("test"),
		//Header: map[string]interface{}{"operator": 331},
		Event: j,
	})

}

// EventStore -
type EventStore struct {
	// 事件編號(traceID)
	EventID string `bson:"eventId"`
	// 事件類型(CRUD)
	EventType int32 `bson:"eventType"`

	// 事件對象
	AggregateID string `bson:"aggregateId"`
	// 對象類別
	AggregateType string `bson:"aggregateType"`

	// 事件資料
	EventData string `bson:"eventData"`

	// 建立時間
	CreatedAt time.Time `bson:"createdAt"`
}

func (e *EventStore) GetEventID() string {
	return e.EventID
}
func (e *EventStore) GetEventType() int32 {
	return int32(e.EventType)
}
func (e *EventStore) GetAggregateID() string {
	return e.AggregateID
}

func (e *EventStore) GetAggregateType() string {
	return "member"
}

func (e *EventStore) GetEventData() string {
	return e.EventData
}

func (e *EventStore) GetCreatedAt() time.Time {
	return e.CreatedAt
}

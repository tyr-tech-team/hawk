package server

import (
	"encoding/json"

	"github.com/tyr-tech-team/hawk/broker"
)

// StopSubscribe - 停止訂閱
func (s *server) StopSubscribe() {
	for _, v := range s.subscriber {
		v.Unsubscribe()
	}
	s.broker.Disconnect()
}

// AddTopic - 新增監聽Topic
func (s *server) AddTopic(topic string, fn func(map[string]interface{}, []byte) error) {
	subscribe, _ := s.broker.Subscribe(
		topic,
		func(e broker.Event) error {
			msg := e.Message()

			if err := fn(msg.Header, msg.Body); err != nil {
				return err
			}
			return nil
		},
		broker.Queue(s.queueName),
	)

	s.subscriber = append(s.subscriber, subscribe)
}

// AddScheduleReply - 新增排程
func (s *server) AddScheduleReply(fn func(map[string]interface{}, *AddScheduleRes) error) {
	subscribe, _ := s.broker.Subscribe(
		ADD_SCHEDULE_REPLY_TOPIC,
		func(e broker.Event) error {
			msg := e.Message()
			res := new(AddScheduleRes)
			json.Unmarshal(msg.Body, res)

			if err := fn(msg.Header, res); err != nil {
				return err
			}
			return nil
		},
		broker.Queue(s.queueName),
	)

	s.subscriber = append(s.subscriber, subscribe)
}

// UpdateScheduleReply - 更新排程
func (s *server) UpdateScheduleReply(fn func(map[string]interface{}, *UpdateScheduleRes) error) {
	subscribe, _ := s.broker.Subscribe(
		UPDATE_SCHEDULE_REPLY_TOPIC,
		func(e broker.Event) error {
			msg := e.Message()
			res := new(UpdateScheduleRes)
			json.Unmarshal(msg.Body, res)

			if err := fn(msg.Header, res); err != nil {
				return err
			}
			return nil
		},
		broker.Queue(s.queueName),
	)

	s.subscriber = append(s.subscriber, subscribe)
}

// CancelScheduleReply - 取消排程
func (s *server) CancelScheduleReply(fn func(map[string]interface{}, *CancelScheduleRes) error) {
	subscribe, _ := s.broker.Subscribe(
		CANCEL_SCHEDULE_REPLY_TOPIC,
		func(e broker.Event) error {
			msg := e.Message()
			res := new(CancelScheduleRes)
			json.Unmarshal(msg.Body, res)

			if err := fn(msg.Header, res); err != nil {
				return err
			}
			return nil
		},
		broker.Queue(s.queueName),
	)

	s.subscriber = append(s.subscriber, subscribe)
}

package event

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type broker struct {
	stream string
	js     nats.JetStream
}

// use broker implements BrokerCast
func (b *broker) Publish(ctx context.Context, subject string, msg Msg) error {
	return b.publish(ctx, subject, msg, false)
}

func (b *broker) PublishWaitAck(ctx context.Context, subject string, msg Msg) error {
	return b.publish(ctx, subject, msg, true)
}

func (b *broker) publish(ctx context.Context, subject string, msg Msg, needAck bool) error {
	ack, err := b.js.PublishAsync(
		subject,
		msg.JSONPayload,
		nats.MsgId(msg.ID),
		nats.ExpectStream(b.stream),
		nats.RetryWait(1*time.Second),
		nats.RetryAttempts(3),
	)

	if err != nil {
		return err
	}

	if needAck {
		select {
		case <-ack.Ok():
			return nil
		case <-ack.Err():
			return err
		}
	}

	return nil
}

func (b *broker) Subscribe(subject string, handler Handler, queue string) error {

	return nil
}

// TODO: 完成此部分
func (b *broker) subscribe(subject string, handler Handler, queue string) error {
	ch := make(chan *nats.Msg, 0)
	if queue != "" {
		b.js.ChanQueueSubscribe(subject, queue, ch, nats.Bind(b.stream))
	}
	return nil
}

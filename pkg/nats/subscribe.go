package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tyr-tech-team/arceus/topic/v1"
	"github.com/tyr-tech-team/hawk/eventbus"
	"github.com/tyr-tech-team/hawk/status"
)

func (cli *jetStreamClient) Subscribe(topic topic.Topic, queue string, handler eventbus.FuncHandler) error {

	var (
		sub *nats.Subscription
		err error
	)

	if _, ok := cli.sublist[topic]; ok {
		return status.CreatedFailed.WithDetail("have the same topic for subscribe").Err()
	}

	sub, err = cli.jc.QueueSubscribe(topic.String(), queue, func(msg *nats.Msg) {
		handler(msg)
	}, nats.BindStream(cli.stream))

	if err != nil {
		return status.ConnectFailed.WithDetail(fmt.Sprintf("connect subscription with topic: %s, failed", topic), err.Error()).Err()
	}

	cli.sublist[topic] = sub

	return nil
}

func (cli *jetStreamClient) CloseSubscription(topic topic.Topic) error {
	sub, ok := cli.sublist[topic]
	if !ok {
		return nil
	}

	if err := sub.Drain(); err != nil {
		return err
	}
	return nil
}

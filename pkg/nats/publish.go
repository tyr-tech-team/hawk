package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/tyr-tech-team/arceus/topic/v1"
	"github.com/tyr-tech-team/hawk/status"
)

func (cli *jetStreamClient) Publish(topic topic.Topic, msg []byte) error {

	_, err := cli.jc.Publish(
		topic.String(),
		msg,
		nats.ExpectStream(cli.stream),
	)

	if err != nil {
		cli.log.Sugar().With("error", err).Errorf("publish with topic[%s] failed", topic.String())
		return status.PublishFailed.WithDetail(err.Error()).Err()
	}

	return nil
}

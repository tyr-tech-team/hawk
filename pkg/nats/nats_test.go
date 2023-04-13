package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/arceus/topic/v1"
	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/eventbus"
	"github.com/tyr-tech-team/hawk/log"
)

func Test_JetStream(t *testing.T) {

	var (
		jc  eventbus.EventBus
		err error
	)
	cfg := config.NatsJetstream{
		Hosts:  "nats://localhost:4223",
		Stream: "test",
		Queue:  "testv1",
	}

	t.Run("init stream", func(t *testing.T) {
		nc, err := nats.Connect(cfg.Hosts)
		assert.NoError(t, err)
		defer nc.Close()

		jc, err := nc.JetStream()
		assert.NoError(t, err)

		_, err = jc.StreamInfo(cfg.Stream)
		if err != nil {
			_, err := jc.AddStream(&nats.StreamConfig{
				Name: cfg.Stream,
				Subjects: []string{
					topic.APIEventALL.String(),
				},
			})
			assert.NoError(t, err)
		}
	})

	t.Run("init queue", func(t *testing.T) {
		nc, err := nats.Connect(cfg.Hosts)
		assert.NoError(t, err)

		defer nc.Close()

		jc, err := nc.JetStream()
		assert.NoError(t, err)

		sub, err := jc.QueueSubscribe(
			topic.APIEventALL.String(),
			cfg.Queue,
			func(msg *nats.Msg) {
				msg.Nak()

			}, nats.BindStream(cfg.Stream))
		assert.NoError(t, err)
		sub.Unsubscribe()

	})

	t.Run("connect jetstream", func(t *testing.T) {
		jc, err = NewJetStream(cfg, log.NewZapLogger("dev"))
		assert.NoError(t, err, "connect nat jetstream failed")
	})

	t.Run("publish message", func(t *testing.T) {
		err := jc.Publish(topic.APIEventALL, []byte("1234"))
		assert.NoError(t, err)
	})

	t.Run("subscribe message", func(t *testing.T) {
		err := jc.Subscribe(topic.APIEventALL, cfg.Queue, func(msg *nats.Msg) {
			assert.Equal(t, string("1234"), string(msg.Data))
		})
		assert.NoError(t, err)
	})

}

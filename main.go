package main

import (
	"fmt"

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
		m := b.Message()

		fmt.Println(string(m.Body))
		fmt.Println(m.Header["operator"])

		return nil

	}, broker.Queue("member"))

	b.Publish("a", &broker.Message{
		Header: map[string]interface{}{},
		Body:   []byte("test"),
	})

}

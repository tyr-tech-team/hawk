package main

import (
	"fmt"

	"github.com/tyr-tech-team/hawk/broker"
	"github.com/tyr-tech-team/hawk/broker/nats"
)

func main() {
	n := nats.NewNats(
		nats.SetURL("172.16.10.70:4222"),
	)

	b := broker.NewBroker(n)

	count := 0
	b.Subscribe("a", func(p broker.Event) error {
		m := p.Message()
		count++
		if count > 900000 {
			fmt.Println(count, string(m.Body))
		}

		return nil
	})

	for i := 0; i < 999999; i++ {
		go b.Publish("a", &broker.Message{
			Header: map[string]interface{}{
				"apple": "mac",
			},
			Body: []byte("adam"),
		})
	}

	select {}
}

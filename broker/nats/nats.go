package nats

import (
	"context"
	"encoding/json"
	"fmt"

	"sync"

	"github.com/nats-io/nats.go"
	natsgo "github.com/nats-io/nats.go"
	"github.com/tyr-tech-team/hawk/broker"
)

const (
	DefaultURL = natsgo.DefaultURL
)

type natsBroker struct {
	sync.RWMutex
	sync.Once
	options broker.Options
	url     string
	client  *natsgo.Conn
}

type publication struct {
	t   string
	err error
	m   *broker.Message
}

type subscriber struct {
	s *nats.Subscription
}

func (p *publication) Topic() string {
	return p.t
}

func (p *publication) Message() *broker.Message {
	return p.m
}

func (p *publication) Error() error {
	return p.err
}

// Topic -
func (s *subscriber) Topic() string {
	return s.s.Subject
}

// Unsubscribe -
func (s *subscriber) Unsubscribe() error {
	return s.s.Unsubscribe()
}

// NatsInstance -
func NewNats(opts ...broker.Option) *natsBroker {
	options := broker.NewOptions(opts...)

	// nats struct
	n := &natsBroker{
		options: options,
		url:     setURL(options),
	}

	// connect
	n.connect()

	return n
}

// Connect -
func (n *natsBroker) connect() {
	n.Lock()
	defer n.Unlock()

	// connect nats
	nc, err := natsgo.Connect(n.url)
	if err != nil {
		panic(err)
	}

	n.client = nc
}

// Publish -
func (n *natsBroker) Publish(topic string, m *broker.Message) error {
	j, _ := json.Marshal(m)

	if err := n.client.Publish(topic, j); err != nil {
		return err
	}

	return nil
}

// Subscribe -
func (n *natsBroker) Subscribe(topic string, h broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	n.RLock()
	defer n.RUnlock()

	opt := broker.SubscribeOptions{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	fn := func(msg *natsgo.Msg) {
		var m broker.Message
		err := json.Unmarshal(msg.Data, &m)

		if err != nil {
			pub := &publication{
				t:   msg.Subject,
				err: err,
				m:   &broker.Message{},
			}

			h(pub)
			return
		}

		pub := &publication{
			t:   msg.Subject,
			err: nil,
			m:   &m,
		}

		h(pub)

	}

	var sub *nats.Subscription
	var err error
	if opt.Queue != "" {
		sub, err = n.client.QueueSubscribe(topic, opt.Queue, fn)
	} else {
		sub, err = n.client.Subscribe(topic, fn)
	}
	if err != nil {
		return nil, err
	}

	return &subscriber{s: sub}, nil
}

// Disconnect -
func (n *natsBroker) Disconnect() error {
	n.Lock()
	defer n.Unlock()

	// drain the connection if specified
	n.client.Drain()
	// close the client connection
	n.client.Close()

	return nil
}

func (n *natsBroker) Address() string {
	if n.client != nil && n.client.IsConnected() {
		return n.client.ConnectedUrl()
	}

	return ""
}

func setURL(opts broker.Options) string {
	u := DefaultURL

	// set URL

	ctxURL, ok := opts.Context.Value(url{}).(string)
	if ok {
		u = ctxURL
	}

	ctxUser, ok := opts.Context.Value(user{}).(string)
	if !ok {
		return u
	}

	ctxPassword, ok := opts.Context.Value(password{}).(string)
	if !ok {
		return u
	}

	return fmt.Sprintf("%s:%s@%s", ctxUser, ctxPassword, u)
}

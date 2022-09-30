// Package natsstreaming -
package natsstreaming

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nuid"
	stan "github.com/nats-io/stan.go"
	"github.com/tyr-tech-team/hawk/broker/v1"
)

const (
	//DefaultURL -
	DefaultURL = stan.DefaultNatsURL

	// DefaultstanClusterID -
	DefaultstanClusterID = "test-cluster"

	// DefaultclientID -
	DefaultclientID = "test-client"
)

type natsstreamingBroker struct {
	sync.RWMutex
	sync.Once
	options       broker.Options
	url           string
	client        stan.Conn
	stanClusterID string
	clientID      string
}

type publication struct {
	t   string
	err error
	m   *broker.Message
}

type subscriber struct {
	s stan.Subscription
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
	return s.Topic()
}

// Unsubscribe -
func (s *subscriber) Unsubscribe() error {
	return s.Unsubscribe()
}

// New -
func New(opts ...broker.Option) broker.Broker {
	n := &natsstreamingBroker{}

	options := broker.NewOptions(opts...)

	n.options = options
	n.stanClusterID = setClusterID(options)
	n.clientID = setClientID(options)
	n.url = setURL(options)

	// connect
	n.connect()

	return n
}

// Connect -
func (n *natsstreamingBroker) connect() {
	n.Lock()
	defer n.Unlock()

	// connect nats
	nc, err := stan.Connect(n.stanClusterID, n.clientID, stan.NatsURL(n.url))
	if err != nil {
		panic(err)
	}

	n.client = nc
}

// Publish -
func (n *natsstreamingBroker) Publish(topic string, m *broker.Message) error {
	j, _ := json.Marshal(m)

	err := n.client.Publish(topic, j)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Subscribe -
func (n *natsstreamingBroker) Subscribe(topic string, h broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	n.RLock()
	defer n.RUnlock()

	opt := broker.SubscribeOptions{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	fn := func(msg *stan.Msg) {
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

	var sub stan.Subscription
	var err error
	if opt.Queue != "" {
		sub, err = n.client.QueueSubscribe(
			topic,
			opt.Queue,
			fn,
			stan.DurableName("durable-name"),
		)
	} else {
		sub, err = n.client.Subscribe(
			topic,
			fn,
			stan.DurableName("durable-name"),
		)
	}
	if err != nil {
		return nil, err
	}

	return &subscriber{s: sub}, nil
}

// Disconnect -
func (n *natsstreamingBroker) Disconnect() error {
	n.Lock()
	defer n.Unlock()

	// close the client connection
	n.client.Close()

	return nil
}

func (n *natsstreamingBroker) Address() string {
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

func setClusterID(opts broker.Options) string {
	u := DefaultstanClusterID

	ctxClusterID, ok := opts.Context.Value(stanClusterID{}).(string)
	if ok {
		if ctxClusterID != "" {
			u = ctxClusterID
		}
	}

	return u
}

func setClientID(opts broker.Options) string {
	u := DefaultclientID

	ctxClientID, ok := opts.Context.Value(clientID{}).(string)
	if ok {
		if ctxClientID != "" {
			u = ctxClientID
		}
	}

	return fmt.Sprintf("%s-%s", u, nuid.Next())
}

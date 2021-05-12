package broker

// Broker -
type Broker interface {
	// Address -
	Address() string

	// Disconnect -
	Disconnect() error
	// Publish -
	Publish(topic string, m *Message) error

	// Subscribe -
	Subscribe(topic string, h Handler, opts ...SubscribeOption) (Subscriber, error)
}

// Handler -
type Handler func(Event) error

// Message -
type Message struct {
	//Header map[string]interface{}
	//Body   []byte
	Event []byte
}

// Event is given to a subscription handler for processing
type Event interface {
	Topic() string
	Message() *Message
	Error() error
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Topic() string
	Unsubscribe() error
}

// NewBroker -
func NewBroker(in Broker) Broker {
	return in
}

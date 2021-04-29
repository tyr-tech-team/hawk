package broker

// Broker -
type Broker interface {
	//Options() Options
	//Address() string
	//Connect() error
	//Disconnect() error
	Publish(topic string, m *Message) error
	Subscribe(topic string, h Handler) (Subscriber, error)
	//String() string
}

type Handler func(Event) error

type Message struct {
	Header map[string]interface{}
	Body   []byte
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

package event

var _ BrokerCast = &broker{}

type broker struct {
	TopicName string
}

// use broker implements BrokerCast
func (b *broker) Publish(topic string, msg interface{}) error {

	return nil
}

func (b *broker) Subscribe(topic string, handler Handler) error {

	return nil
}

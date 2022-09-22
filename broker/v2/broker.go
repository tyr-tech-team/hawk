// Package broker -
package broker

import (
	"errors"
	"reflect"
	"sync"

	jsoniter "github.com/json-iterator/go"

	micro "go-micro.dev/v4/broker"

	"github.com/go-micro/plugins/v4/broker/stan"
)

// emptyType - 空白類型
var emptyType = reflect.TypeOf(&micro.Message{})

// Broker -
type Broker interface {
	// 監聽
	Sub(topic string, event Handler, quene string)
	// 推播
	Pub(topic string, message interface{}) error
	// 關閉
	Close()
}

// Handler -
type Handler interface{}

// broker -
type broker struct {
	// 互斥鎖
	mu *sync.Mutex
	// 編碼器
	coder jsoniter.API
	// 代理
	mBroker micro.Broker
	// 監聽列表
	subList map[string]micro.Subscriber
}

// New -
func New(host, clusterID string) Broker {
	mBroker := stan.NewBroker(
		micro.Addrs(host),
		stan.ClusterID(clusterID),
	)

	b := &broker{
		mu:      &sync.Mutex{},
		coder:   jsoniter.ConfigCompatibleWithStandardLibrary,
		mBroker: mBroker,
		subList: make(map[string]micro.Subscriber),
	}

	// connect
	if err := b.mBroker.Connect(); err != nil {
		// do log
	}

	return b
}

// Sub -
func (b *broker) Sub(topic string, event Handler, quene string) {
	// event is required
	if event == nil {
		return
	}

	// get arguments type and num
	argType, numArgs, err := argInfo(event)
	if argType == nil || err != nil {
		return
	}

	// reflect event to Value
	eventValue := reflect.ValueOf(event)

	fn := func(e micro.Event) error {
		var oV []reflect.Value

		switch argType {
		case emptyType:
			oV = []reflect.Value{reflect.ValueOf(e)}
		default:
			var oPtr reflect.Value
			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.New(argType)
			} else {
				oPtr = reflect.New(argType.Elem())
			}

			// decoder
			if err := b.coder.Unmarshal(e.Message().Body, oPtr.Interface()); err != nil {
				// do log
			}

			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.Indirect(oPtr)
			}

			// Callback
			switch numArgs {
			case 1:
				oV = []reflect.Value{oPtr}
			}
		}

		eventValue.Call(oV)

		return nil
	}

	sub, _ := b.mBroker.Subscribe(topic, fn, micro.Queue(quene))

	// lock
	b.mu.Lock()

	defer b.mu.Unlock()

	// add map list
	b.subList[topic] = sub
}

// Pub -
func (b *broker) Pub(topic string, v interface{}) error {
	data, err := b.coder.Marshal(v)
	if err != nil {
		return err
	}

	e := &micro.Message{
		Body: data,
	}

	return b.mBroker.Publish(topic, e)
}

// Close -
func (b *broker) Close() {
	b.mu.Lock()

	defer b.mu.Unlock()

	for _, v := range b.subList {
		v.Unsubscribe()
	}
}

// argInfo -
func argInfo(cb Handler) (reflect.Type, int, error) {
	cbType := reflect.TypeOf(cb)
	if cbType.Kind() != reflect.Func {
		return nil, 0, errors.New("Handler need to be a func")
	}

	numArgs := cbType.NumIn()
	if numArgs == 0 {
		return nil, numArgs, nil
	}

	return cbType.In(numArgs - 1), numArgs, nil
}

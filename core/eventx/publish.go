package eventx

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
)

type DataChannel chan *DataEvent

type DataChannelSlice []DataChannel

type DataEvent struct {
	Data  interface{}
	Topic string
}

type EventBus struct {
	Subscribers map[string]DataChannelSlice
	lk          sync.RWMutex
	Ctx         context.Context
	Cancel      context.CancelFunc
}

func (eb *EventBus) Cancl(topic string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Publish recover err")
			debug.PrintStack()
		}
	}()
	eb.lk.Lock()
	defer eb.lk.Lock()
	if chans, found := eb.Subscribers[topic]; found {
		for _, channel := range chans {
			close(channel)
		}
		delete(eb.Subscribers, topic)
	}
}

func (eb *EventBus) Subscribe(topic string) DataChannel {
	eb.lk.Lock()
	defer eb.lk.Unlock()
	ch := make(DataChannel)
	if prev, ok := eb.Subscribers[topic]; ok {
		eb.Subscribers[topic] = append(prev, ch)
	} else {
		eb.Subscribers[topic] = append([]DataChannel{}, ch)
	}
	return ch
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Publish recover err")
			debug.PrintStack()
		}
	}()
	eb.lk.RLock()
	defer eb.lk.RUnlock()
	if chans, found := eb.Subscribers[topic]; found {
		channels := append([]DataChannel{}, chans...)
		go func(data *DataEvent, dataChannelSlice DataChannelSlice) {
			for _, ch := range dataChannelSlice {
				ch <- data
			}
		}(&DataEvent{Data: data, Topic: topic}, channels)
	}
}

package kmq

import (
	"reflect"
	"sync"
)

// Subscriber used to get message data form chan
type Subscriber struct {
	Typ   reflect.Type
	Topic string
	Key   string
	//Data chan []byte
	Func func(data []byte)
}

func defaultSubFunc(data []byte) { /* do nothing */ }

//subscribers hold all subscribers
var subscribers map[string][]*Subscriber
var lock sync.RWMutex

func init() {
	subscribers = make(map[string][]*Subscriber)
}

// SubEvent subscribe message by given interface
// - key is type name of the input interface
func SubEvent(input interface{}, topic string, callback func(data []byte)) *Subscriber {
	lock.Lock()
	defer lock.Unlock()

	refVal := reflect.ValueOf(input)
	key := reflect.Indirect(refVal).Type().Name()
	sub := &Subscriber{
		Typ:   reflect.TypeOf(input),
		Topic: topic,
		Key:   key,
		Func:  callback,
		//Data: make(chan []byte),
	}
	if callback == nil {
		sub.Func = defaultSubFunc
	}
	if subs, exist := subscribers[key]; !exist {
		newSubs := []*Subscriber{}
		subscribers[key] = append(newSubs, sub)
	} else {
		subscribers[key] = append(subs, sub)
	}
	return sub
}

// UnsubEvent unsubscribe message by given interface
func UnsubEvent(toDel *Subscriber) {
	// todo
}

// fireEvent despatch message data for each subscriber
// data shoule be decode by reflect.Type in Subscriber
func fireEvent(topic, key string, data []byte, c *Consumer) {
	lock.RLock()
	defer lock.RUnlock()
	if subs, exist := subscribers[key]; exist {
		for _, sub := range subs {
			if sub != nil {
				if sub.Topic == c.topic {
					switch c.subMode {
					case SubMode_Sync:
						sub.Func(data)
					case SubMode_Asnyc:
						go func() {
							// catch panic, to protect main goroutine
							defer func() {
								if p := recover(); p != nil {
									// todo
									c.loger.Error("event callbace painc! key:%s, subscriber:%#v", key, *sub)
								}
							}()

							// callback
							sub.Func(data)
						}()

					default:
						sub.Func(data)
					}
				}
			}
		}
	}
}

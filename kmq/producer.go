package kmq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/Shopify/sarama"
)

// Producer - used async producer in sarama
type Producer struct {
	asyncProducer sarama.AsyncProducer
	loger         Logger

	wg sync.WaitGroup
}

// NewProducer create an new producer with address and loger, then start loop
func NewProducer(address []string, l Logger) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_0_0_0

	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		return nil, fmt.Errorf("Create producer failed! error:", err)
	}

	asnyncp := &Producer{
		asyncProducer: p,
		loger:         l,
	}

	asnyncp.loop()
	asnyncp.loger.Info("Start producer success! ip:", address)
	asnyncp.wg.Add(2)
	return asnyncp, nil
}

// SendMsg send message to kafka broker
// - value will be encode by json format
// - error is only return the result of local func call, the actually result send to broker can be get from producer's chan
func (p *Producer) SendMsg(topic string, value interface{}) error {
	refVal := reflect.ValueOf(value)
	key := reflect.Indirect(refVal).Type().Name()
	byteValue, err := json.Marshal(value)
	if err != nil {
		p.loger.Error("[Producer] SendMsg error! Marshal message failed! topic:%s, key:%s, value:%v", topic, key, value)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(byteValue),
	}

	p.asyncProducer.Input() <- msg

	return nil
}

// Stop shutdown the producer
// - continue to read from successes and errors channels
func (p *Producer) Stop() {
	p.asyncProducer.AsyncClose()

	p.wg.Wait()
	p.loger.Info("[Producer] Stoped!")
}

// loop create for loop for get message status return from broker
func (p *Producer) loop() {
	// handle success chan
	go func() {
		defer p.wg.Done()
		for s := range p.asyncProducer.Successes() {
			p.loger.Debug("[Producer] Send message sucessed! Topic:%s, partitions:%d, offset:%d, timestamp:%s", s.Topic, s.Partition, s.Offset, s.Timestamp.String())
		}
	}()

	// handle error chan
	go func() {
		defer p.wg.Done()
		for err := range p.asyncProducer.Errors() {
			p.loger.Error("[Producer] Send message error!", err.Err)
		}
	}()
}

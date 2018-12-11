package kmq

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	SubMode_Sync  = 0
	SubMode_Asnyc = 1
)

//====================================================
var client sarama.Consumer

// Consumer get message from kafka broker
type Consumer struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
	topic             string
	loger             Logger
	offset            Offset
	subMode           int
	running           bool // prevent duplicated Start call

	wg sync.WaitGroup
}

// NewConsumer create an new consumer with address and loger
func NewConsumer(address []string, l Logger, topic string, partition int32, offset Offset, mode int) (*Consumer, error) {
	if offset == nil {
		return nil, fmt.Errorf("Create consumer failed, offset is nil! ")
	}

	if mode != SubMode_Sync && mode != SubMode_Asnyc {
		return nil, fmt.Errorf("Create consumer failed, input mode is unexpected! please input SubMode_Sync or SubMode_Asnyc")
	}

	if client == nil {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Version = sarama.V2_0_0_0
		var err error
		client, err = sarama.NewConsumer(address, config)
		if err != nil {
			return nil, fmt.Errorf("Create consumer failed! error:", err)
		}
	}

	loadOffset, err := offset.Load()
	if err != nil {
		return nil, fmt.Errorf("load offset failed! error:", err)
	}

	partitionConsume, err := client.ConsumePartition(topic, partition, loadOffset)
	if err != nil {
		return nil, fmt.Errorf("Create partition consumer failed! error:", err)
	}

	newConsumer := &Consumer{
		consumer:          client,
		partitionConsumer: partitionConsume,
		loger:             l,
		offset:            offset,
		subMode:           mode,
		topic:             topic,
	}

	return newConsumer, nil
}

// Start start handle message from kafka
// - !notice: before call Start, you should already subscribed the events that you want to handle
func (c *Consumer) Start() {
	if !c.running {
		c.loop()
		c.loger.Info("Start consumer success!")
		c.wg.Add(2)
		c.running = true
	}
}

// Stop shutdown the Consumer
// - continue to read from successes and errors channels
func (c *Consumer) Stop() {
	c.partitionConsumer.AsyncClose()
	// wait loop return
	c.wg.Wait()
	c.running = false
	c.loger.Info("[Consumer] Stoped!")
}

// loop create for loop for get message status return from broker
func (c *Consumer) loop() {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				// todo
				// recover
			}
		}()

		defer c.wg.Done()
		for msg := range c.partitionConsumer.Messages() {
			c.loger.Debug("--->[Consumer] get new message, topic:%s, offset:%d, key:%s", msg.Topic, msg.Offset, string(msg.Key))
			err := c.offset.Save(msg.Offset)
			if err != nil {
				c.loger.Error("!!! [Consumer] save offset %d failed: %v", msg.Offset, err)
			}

			fireEvent(msg.Topic, string(msg.Key), msg.Value, c)
		}
	}()

	go func() {
		defer c.wg.Done()
		for err := range c.partitionConsumer.Errors() {
			c.loger.Error("[Consumer] get error at topic:%s, partition:%d, reson:%s", err.Topic, err.Partition, err.Error())
		}
	}()
}

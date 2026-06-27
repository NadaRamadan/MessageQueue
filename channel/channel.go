package channel

import (
	"nsq-clone/consumer"
	"nsq-clone/message"
	"sync"
)

type Channel struct {
	name      string
	messages  chan *message.Message
	consumers []*consumer.Consumer
}

func New(name string) *Channel {
	return &Channel{
		name:      name,
		messages:  make(chan *message.Message, 100),
		consumers: make([]*consumer.Consumer, 0),
	}
}

func (c *Channel) Enqueue(msg *message.Message) {
	c.messages <- msg
}

func (c *Channel) AddConsumer(cons *consumer.Consumer) {
	c.consumers = append(c.consumers, cons)
}

func (c *Channel) StartConsumers(wg *sync.WaitGroup) {

	for _, cons := range c.consumers {

		wg.Add(1)

		go func(consumer *consumer.Consumer) {

			defer wg.Done()

			for msg := range c.messages {
				println("Consumer", consumer.Id, "processed", msg.Content)
			}

		}(cons)
	}
}
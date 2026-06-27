package topic

import (
	"sync"

	"encoding/json"
	"nsq-clone/channel"
	"nsq-clone/message"
	"os"
)

type Topic struct {
	name     string
	channels map[string]*channel.Channel
	messages chan *message.Message
	mu       sync.RWMutex
}

func New(name string) *Topic {
	return &Topic{
		name:     name,
		channels: make(map[string]*channel.Channel),
		messages: make(chan *message.Message, 100),
	}
}

func (t *Topic) AddChannel(name string) *channel.Channel {
	t.mu.Lock()
	defer t.mu.Unlock()

	ch := channel.New(name)
	t.channels[name] = ch
	return ch
}

func (t *Topic) Broadcast(msg *message.Message) {
	t.messages <- msg

	go func() {
		t.mu.RLock()
		defer t.mu.RUnlock()

		for _, ch := range t.channels {
			ch.Enqueue(msg)
		}
	}()
}
func (t *Topic) PersistMessage(msg *message.Message) error {

	file, err := os.OpenFile(
		"messages.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	data, _ := json.Marshal(msg)

	_, err = file.Write(append(data, '\n'))

	return err
}

package message

type Message struct {
	id      string
	Content string
}

func New(id, content string) *Message {
	return &Message{
		id:      id,
		Content: content,
	}
}





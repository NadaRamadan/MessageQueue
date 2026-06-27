package consumer

type Consumer struct {
	Id string
}

func New(id string) *Consumer {
	return &Consumer{Id: id}
}

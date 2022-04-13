package model

type client struct {
	Transport chan interface{}
	name      string
}

func NewClient(name string) *client {
	return &client{
		name:      name,
		Transport: make(chan interface{}),
	}
}

func (C *client) Consume(message interface{}) {
	C.Transport <- message
}

func (C *client) GetName() string {
	return C.name
}

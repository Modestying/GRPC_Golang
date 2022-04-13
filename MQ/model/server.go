package model

import (
	"fmt"
	"sync"
)

type MQ struct {
	clients map[string]*client
	sync.RWMutex
}

func NewMQ() *MQ {
	return &MQ{
		clients: make(map[string]*client),
	}
}

func (S *MQ) AddClient(cli *client) {
	if _, exist := S.clients[cli.name]; !exist {
		S.Lock()
		S.clients[cli.name] = cli
		S.Unlock()
	}
}

func (S *MQ) UnSubscribe(name string) {
	if _, exist := S.clients[name]; exist {
		S.Lock()
		delete(S.clients, name)
		S.Unlock()
	}
}

func (S *MQ) Notify(message interface{}) {
	if len(S.clients) != 0 {
		for _, v := range S.clients {
			fmt.Printf("%v ", v.GetName())
			v.Consume(message)
		}
		fmt.Println()
	}
}

package channel

import (
	"fmt"
	"sync"
)

type Broker struct {
	mutex sync.Mutex

	chans []chan Msg
}

func (b *Broker) Send(m Msg) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, c := range b.chans {
		select {
		case c <- m:
		default:
			return fmt.Errorf("channel is full") // 如果channel满了，就返回错误
		}
	}
	return nil
}

func (b *Broker) Subscribe(cap int) (<-chan Msg, error) {
	res := make(chan Msg, cap)

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.chans = append(b.chans, res)
	return res, nil
}

type Msg struct {
	Content string
}

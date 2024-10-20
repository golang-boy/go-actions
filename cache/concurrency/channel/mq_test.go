package channel

import (
	"sync"
	"testing"
	"time"
)

func TestMq(t *testing.T) {
	b := &Broker{}

	go func() {

		for {
			err := b.Send(Msg{Content: time.Now().String()})
			if err != nil {
				t.Log(err)
				return
			}

			time.Sleep(1 * time.Millisecond)
		}

	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {

		wg.Add(i)
		go func(i int) {
			t.Logf("消费者: %d", i)

			defer wg.Done()
			msgs, err := b.Subscribe(100)
			if err != nil {
				t.Log(err)
				return
			}

			for msg := range msgs {
				t.Log(msg.Content)
			}
		}(i)
	}

	wg.Wait()
}

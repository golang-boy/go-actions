package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	var result int64 = 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()

			atomic.AddInt64(&result, i)
		}(int64(i))
	}
	wg.Wait()
	fmt.Printf("result: %d\n", result)
}

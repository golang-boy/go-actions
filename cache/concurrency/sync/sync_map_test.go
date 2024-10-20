package sync

import (
	"testing"
	"time"
)

func TestSafeMap(t *testing.T) {

	s := &SafeMap[string, string]{
		data: make(map[string]string),
	}

	go func() {
		val, ok := s.LoadOrStore("key", "value1")
		t.Log("goroutine1: ", val, ok)
	}()

	go func() {
		val, ok := s.LoadOrStore("key", "value2")
		t.Log("goroutine2: ", val, ok)
	}()

	// 期望谁先设置，谁先获取到

	time.Sleep(time.Second)
}

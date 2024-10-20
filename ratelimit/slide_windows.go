package ratelimit

import (
	"container/list"
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
)

type SlideWindows struct {
	interval time.Duration
	queue    *list.List

	rate int
}

func (t *SlideWindows) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		now := time.Now().UnixNano()
		boundary := now - int64(t.interval)
		timestamp := t.queue.Front()

		for timestamp.Value.(int64) < boundary {
			t.queue.Remove(t.queue.Front())
			timestamp = t.queue.Front()
		}

		if t.queue.Len() >= t.rate {
			return nil, errors.New("rate limit")
		}

		resp, err = handler(ctx, req)
		t.queue.PushBack(now)

		return
	}
}

package ratelimit

import (
	"context"
	"time"

	"errors"

	"google.golang.org/grpc"
)

type FixWindows struct {
	timestamp int64
	interval  time.Duration

	rate int64
	cnt  int64
}

func NewFixWindows(rate int64, interval time.Duration) *FixWindows {

	return &FixWindows{
		timestamp: time.Now().UnixNano(),
		interval:  interval,
		rate:      rate,
		cnt:       0,
	}

}

func (t *FixWindows) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		cur := time.Now().UnixNano()
		if t.timestamp+int64(t.interval) < cur {
			t.timestamp = cur
			t.cnt = 0
		}

		if t.rate <= t.cnt {
			err = errors.New("rate limit exceeded")
			return
		}
		t.cnt++
		resp, err = handler(ctx, req)
		return
	}
}

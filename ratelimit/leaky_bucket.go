package ratelimit

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type LeakyBucket struct {
	producer *time.Ticker
}

func NewLeakyBucket(interval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		producer: time.NewTicker(interval),
	}
}

func (t *LeakyBucket) BuildServerInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		select {
		case <-t.producer.C:
			resp, err = handler(ctx, req)
		case <-ctx.Done():
			err = ctx.Err()
			return
		}

		return
	}
}

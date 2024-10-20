package ratelimit

import (
	"context"

	"google.golang.org/grpc"
)

type TokenBucketLimiter struct {
	tokens chan struct{}
}

func NewTokenBucketLimiter(rate int) *TokenBucketLimiter {
	t := &TokenBucketLimiter{
		tokens: make(chan struct{}, rate),
	}

	for i := 0; i < rate; i++ {
		t.tokens <- struct{}{}
	}
	return t
}

func (t *TokenBucketLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		case <-t.tokens:
			resp, err = handler(ctx, req)
		}
		return
	}
}

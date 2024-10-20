package ratelimit

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

//go:embed lua/slide_window.lua
var luaSliedWindow string

type RedisSliedWindow struct {
	client   redis.Cmdable
	interval time.Duration
	rate     int64
}

func NewRedisSliedWindow(client redis.Cmdable, service string, interval int64, rate int64) *RedisSliedWindow {
	return &RedisSliedWindow{client: client}
}

func (t *RedisSliedWindow) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if allow, err1 := t.allow(ctx, info.FullMethod); !allow {
			if err1 != nil {
				return nil, err1
			}
			err = errors.New("rate limit")
			return
		}
		resp, err = handler(ctx, req)
		return
	}
}

func (t *RedisSliedWindow) allow(ctx context.Context, service string) (bool, error) {
	return t.client.Eval(ctx, luaSliedWindow, []string{service}, t.interval.Milliseconds(), t.rate, time.Now().UnixMilli()).Bool()
}

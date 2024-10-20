package ratelimit

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

//go:embed lua/fix_window.lua
var luaFixWindow string

type RedisFixWindow struct {
	client redis.Cmdable
}

func NewRedisFixWindow(client redis.Cmdable, service string, interval int64, rate int64) *RedisFixWindow {
	return &RedisFixWindow{client: client}
}

func (t *RedisFixWindow) BuildServerInterceptor() grpc.UnaryServerInterceptor {
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

func (t *RedisFixWindow) allow(ctx context.Context, service string) (bool, error) {
	return t.client.Eval(ctx, luaFixWindow, []string{service}, 1).Bool()
}

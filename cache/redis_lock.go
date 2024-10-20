package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client redis.Cmdable
}

type Lock struct {
}

func (c *Client) TryLock(ctx context.Context, key string, ttl time.Duration) (*Lock, error) {

	ok, err := c.client.SetNX(ctx, key, "locked", ttl).Result()
	if err != nil {
		return nil, err
	}

	if !ok {
		// log.Printf("lock %s failed", key)
		return &Lock{}, nil
	}

	return &Lock{}, nil
}

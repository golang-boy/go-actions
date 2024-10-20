package cache

import (
	"context"
	"time"
)

type ReadThroughCache struct {
	Cache

	LoadFunc func(ctx context.Context, key string) (any, error)
}

func (c *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	// Check if the value is in the cache
	if value, err := c.Cache.Get(ctx, key); err == errKeyNotFound {

		defaultExpiration := 100 * time.Second

		// If not, load the value from the data source
		val, err := c.LoadFunc(ctx, key)
		if err == nil {
			err = c.Cache.Set(ctx, key, val, defaultExpiration)
		}

		return value, nil
	}

	return nil, nil

}

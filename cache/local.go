package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	errKeyNotFound = errors.New("key not found")
)

type BuildInMapCache struct {
	data  map[string]*item
	mutex sync.Mutex
}

type item struct {
	value    any
	deadline time.Time
}

func NewBuildInMapCache() *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item, 100),
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)

		for t := range ticker.C {
			res.mutex.Lock()
			for k, v := range res.data {
				if !v.deadline.IsZero() && v.deadline.Before(t) {
					delete(res.data, k)
				}
			}
			res.mutex.Unlock()
		}
	}()

	return res
}

func (b *BuildInMapCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var dl time.Time

	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}

	b.data[key] = &item{
		value:    value,
		deadline: dl,
	}

	if expiration > 0 {
		time.AfterFunc(expiration, func() {
			b.mutex.Lock()
			defer b.mutex.Unlock()

			val, ok := b.data[key]
			if ok && !val.deadline.IsZero() && val.deadline.Before(time.Now()) {
				delete(b.data, key)
			}
		})

	}
	return nil
}

func (b *BuildInMapCache) Get(ctx context.Context, key string) (any, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	res, ok := b.data[key]
	if !ok {
		return nil, errKeyNotFound
	}
	return res, nil
}

func (b *BuildInMapCache) Delete(ctx context.Context, key string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	delete(b.data, key)
	return nil
}

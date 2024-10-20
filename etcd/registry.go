package etcd

import (
	"context"

	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/clientv3"
)

type Registry struct {
	c    *clientv3.Client
	sess *concurrency.Session
}

func NewRegistry(c *clientv3.Client) (*Registry, error) {
	sess, err := concurrency.NewSession(c)
	if err != nil {
		panic(err)
	}
	return &Registry{
		c:    c,
		sess: sess,
	}, err
}

func (r *Registry) Register(ctx context.Context, i registry)

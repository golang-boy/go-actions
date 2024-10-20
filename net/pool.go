package net

import (
	"context"
	"errors"
	"net"
	"time"
)

type Pool struct {
	idlesConns chan *idleConn

	reqQueue []connReq

	maxCnt int

	cnt int

	maxIdleTime time.Duration

	initCnt int
	factory func() (net.Conn, error)
}

type idleConn struct {
	c              net.Conn
	lastActiveTime time.Time
}

func NewPool(initCnt int, maxidleCnt int, maxCnt int, maxIdleTime time.Duration, factory func() (net.Conn, error)) (*Pool, error) {

	idlesConn := make(chan *idleConn, maxidleCnt)

	if initCnt > maxidleCnt {
		return nil, errors.New("initCnt should be less than maxidleCnt")
	}

	for i := 0; i < initCnt; i++ {
		c, err := factory()
		if err != nil {
			return nil, err
		}
		idlesConn <- &idleConn{c: c, lastActiveTime: time.Now()}
	}

	res := &Pool{
		idlesConns:  make(chan *idleConn, maxidleCnt),
		maxCnt:      maxCnt,
		maxIdleTime: maxIdleTime,
		initCnt:     initCnt,
		factory:     factory,
		reqQueue:    make([]connReq, 0),
		cnt:         0,
	}

	return res, nil
}

func (p *Pool) Get(ctx context.Context) (net.Conn, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	for {
		select {
		case c := <-p.idlesConns:
			if time.Since(c.lastActiveTime) > p.maxIdleTime {
				c.c.Close()
				continue
			}
		default:
			// p.lock.Lock()
			// if p.cnt >= p.maxCnt {
			// 	req := connReq{connChan: make(chan net.Conn, 1)}

			// }
		}
	}
}

type connReq struct {
	connChan chan net.Conn
}

package sync

import "sync"

type Biz struct {
	once sync.Once
}

func (m *Biz) Init() {
	m.once.Do(func() {
		// 初始化操作
	})
}

type Busi interface {
	Do()
}

type singleton struct {
}

func (s *singleton) Do() {
}

var s *singleton
var singletonOnce sync.Once

func GetSingleton() Busi {
	singletonOnce.Do(func() {
		s = &singleton{}
	})
	return s
}

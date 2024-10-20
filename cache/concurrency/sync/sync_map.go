package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) Set(key K, value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *SafeMap[K, V]) Get(key K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *SafeMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {

	s.mutex.RLock()

	res, ok := s.data[key]
	s.mutex.RUnlock()
	if ok {
		return res, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	res, ok = s.data[key]
	if ok {
		return res, true
	}

	s.data[key] = value
	return value, false
}

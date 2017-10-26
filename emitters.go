package xray

import "sync"

// NewSyncEmitter builds and returns new synchronous event emitter
func NewSyncEmitter() EventEmitter {
	return &syncEmitter{}
}

type syncEmitter struct {
	m         sync.Mutex
	listeners []func(...Event)
}

func (s *syncEmitter) Emit(events ...Event) {
	if len(events) > 0 {
		for _, listener := range s.listeners {
			listener(events...)
		}
	}
}

func (s *syncEmitter) On(c func(...Event)) {
	s.m.Lock()
	defer s.m.Unlock()

	s.listeners = append(s.listeners, c)
}

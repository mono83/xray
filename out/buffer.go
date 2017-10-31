package out

import (
	"github.com/mono83/xray"
	"sync"
	"time"
)

// BufferOneSecond builds wrapper over events handler, that buffers all events it receives
// and then flushes them every second
func BufferOneSecond(target xray.Handler) xray.Handler {
	return Buffer(target, time.Second)
}

// Buffer builds wrapper over events handler, that buffers all events it receives
// and then flushes them on regular basis
func Buffer(target xray.Handler, interval time.Duration) xray.Handler {
	ch := &buffered{
		interval: interval,
		target:   target,
	}

	go ch.wait()
	return ch.receive
}

type buffered struct {
	sync.Mutex
	buffer []xray.Event

	interval time.Duration
	target   xray.Handler
}

// Flush method flushes events to target
func (b *buffered) Flush() {
	b.Lock()
	local := b.buffer
	if len(local) > 0 {
		b.buffer = []xray.Event{}
	}
	b.Unlock()
	if len(local) > 0 {
		b.target(local...)
	}
}

func (b *buffered) wait() {
	for {
		time.Sleep(b.interval)
		b.Flush()
	}
}

func (b *buffered) receive(events ...xray.Event) {
	if len(events) > 0 {
		b.Lock()
		b.buffer = append(b.buffer, events...)
		b.Unlock()
	}
}

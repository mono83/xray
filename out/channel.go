package out

import "github.com/mono83/xray"

// Channel builds wrapper over events handler, that works in asynchronous mode
// using delivery goroutine and channel
func Channel(target xray.Handler) func(...xray.Event) {
	ch := &channeled{
		ch:     make(chan xray.Event),
		target: target,
	}

	go ch.wait()
	return ch.receive
}

type channeled struct {
	ch     chan xray.Event
	target xray.Handler
}

func (c *channeled) wait() {
	for event := range c.ch {
		c.target(event)
		xray.Waiter.Done()
	}
}

func (c *channeled) receive(events ...xray.Event) {
	xray.Waiter.Add(len(events))
	for _, e := range events {
		c.ch <- e
	}
}

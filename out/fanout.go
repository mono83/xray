package out

import (
	"github.com/mono83/xray"
)

// FanOut combines multiple handlers into one.
// All combined handlers will receive all events
func FanOut(handlers ...xray.Handler) xray.Handler {
	if len(handlers) == 0 {
		return nil
	}
	if len(handlers) == 1 {
		return handlers[0]
	}

	foh := fanoutHandlers(handlers)

	return foh.handle
}

type fanoutHandlers []xray.Handler

func (f fanoutHandlers) handle(events ...xray.Event) {
	for _, handle := range f {
		handle(events...)
	}
}

package out

import (
	"github.com/mono83/xray"
)

// Splitter returns handler, that will split incoming events slice into smaller slices with configured size
func Splitter(target xray.Handler, size int) xray.Handler {
	return func(events ...xray.Event) {
		l := len(events)
		if l == 0 {
			return
		} else if l <= size {
			target(events...)
		} else {
			chunks := l / size
			if chunks*size < l {
				chunks++
			}
			var from, to int
			for i := 0; i < chunks; i++ {
				from = i * size
				to = (i + 1) * size
				if to > l {
					to = l
				}
				target(events[from:to]...)
			}
		}
	}
}

// SplitterOne returns handler, that will split incoming events one by one
func SplitterOne(target func(xray.Event)) xray.Handler {
	return func(events ...xray.Event) {
		if l := len(events); l == 0 {
			return
		}

		for _, event := range events {
			target(event)
		}
	}
}

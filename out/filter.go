package out

import "github.com/mono83/xray"

// Filter provides handler filtering adapter
func Filter(target xray.Handler, predicate func(event xray.Event) bool) xray.Handler {
	return func(events ...xray.Event) {
		l := len(events)
		if l == 0 {
			return
		} else if l == 1 {
			if predicate(events[0]) {
				target(events[0])
			}
		} else {
			delivery := []xray.Event{}
			for _, e := range events {
				if predicate(e) {
					delivery = append(delivery, e)
				}
			}

			if len(delivery) > 0 {
				target(delivery...)
			}
		}
	}
}

// FilterMetrics returns filtering adapter, that will forward to target only metric events
func FilterMetrics(target xray.Handler) xray.Handler {
	return Filter(target, func(event xray.Event) bool {
		if event == nil {
			return false
		}
		_, ok := event.(xray.MetricsEvent)
		return ok
	})
}

// FilterLogs returns filtering adapter, that will forward to target only log events with minimal
// configured level
func FilterLogs(target xray.Handler, level xray.Level) xray.Handler {
	return Filter(target, func(event xray.Event) bool {
		if event == nil {
			return false
		}
		l, ok := event.(xray.LogEvent)
		return ok && l.GetLevel() >= level
	})
}

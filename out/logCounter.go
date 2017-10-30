package out

import (
	"github.com/mono83/xray"
)

// LogCounter returns handler, that will count all logging events and send counters as metrics
func LogCounter(r xray.Ray) xray.Handler {
	return func(events ...xray.Event) {
		for _, e := range events {
			if e == nil {
				continue
			}

			if l, ok := e.(xray.LogEvent); ok {
				switch l.GetLevel() {
				case xray.TRACE:
					r.Inc("trace")
				case xray.DEBUG:
					r.Inc("debug")
				case xray.INFO:
					r.Inc("info")
				case xray.WARNING:
					r.Inc("warning")
				case xray.ERROR:
					r.Inc("error")
				case xray.ALERT:
					r.Inc("alert")
				case xray.CRITICAL:
					r.Inc("critical")
				}
			}
		}
	}
}

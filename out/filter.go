package out

import "github.com/mono83/xray"

// Filter provides filtering adapter
func Filter(target func(...xray.Event), predicate func(event xray.Event) bool) func(...xray.Event) {
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

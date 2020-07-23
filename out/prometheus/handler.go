package prometheus

import (
	"github.com/mono83/xray"
)

// Handle handles logging/metrics events
func (e *Exporter) Handle(events ...xray.Event) {
	e.mutex.Lock()
	for _, event := range events {
		if m, ok := event.(xray.MetricsEvent); ok {
			e.cntHandled++
			// Calculating metric key
			key := xray.MetricGroupingKey(m, e.filter)

			switch m.GetType() {
			case xray.INCREMENT:
				v, ok := e.counters[key]
				if !ok {
					v = &value{args: e.filter(m.Args()), name: m.GetKey()}
					e.counters[key] = v
				}
				v.value += m.GetValue()
			case xray.GAUGE:
				v, ok := e.gauges[key]
				if !ok {
					v = &value{args: e.filter(m.Args()), name: m.GetKey()}
					e.gauges[key] = v
				}
				v.value = m.GetValue()
			case xray.DURATION:
				v, ok := e.histogram[key]
				if !ok {
					v = &timeValue{
						buckets: make([]int64, len(e.buckets)),
						counts:  0,
						sum:     0,
						name:    m.GetKey(),
						args:    e.filter(m.Args()),
					}
					e.histogram[key] = v
				}

				nano := m.GetValue()

				v.counts++
				v.sum += nano
				for i, j := range e.buckets {
					if nano <= j {
						v.buckets[i]++
						break
					}
				}
			}
		}
	}
	e.mutex.Unlock()
}

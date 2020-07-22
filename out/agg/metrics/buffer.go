package metrics

import (
	"sync"
	"time"

	"github.com/mono83/xray"
)

// NewBuffer builds and returns new metrics buffer
// No automatic flush, one must invoke FlushTo manually
func NewBuffer(filter xray.ArgFilter, percentiles ...int) *Buffer {
	if filter == nil {
		filter = func(args []xray.Arg) []xray.Arg {
			return args
		}
	}
	return &Buffer{
		percentiles: percentiles,
		ArgFilter:   filter,
		inc:         NewIncrementalHolder(),
		gauge:       NewGaugeHolder(),
		delta:       NewDeltaHolder(percentiles...),
		definitions: map[string]definition{},
	}
}

// NewBufferFunc builds buffer and returns handler func, compatible with xray.Handler
// This buffer will automatically flush all data in fixed rate, configured as interval
func NewBufferFunc(interval time.Duration, target xray.Handler, filter xray.ArgFilter, percentiles ...int) xray.Handler {
	buf := NewBuffer(filter, percentiles...)
	go func() {
		for {
			time.Sleep(interval)
			buf.FlushTo(target)
		}
	}()

	return buf.Handle
}

// Buffer is metrics buffer, that accumulates metric data and then flushes them
// to provided target.
// This container is thread-safe
type Buffer struct {
	sync.Mutex
	ArgFilter         xray.ArgFilter
	inc, gauge, delta GroupingBuffer
	percentiles       []int
	definitions       map[string]definition
}

// Handle method is xray.Handler implementation
func (b *Buffer) Handle(events ...xray.Event) {
	for _, event := range events {
		if m, ok := event.(xray.MetricsEvent); ok {
			b.Add(m)
		}
	}
}

// Add method adds new metrics event to buffer
// Thread safe
func (b *Buffer) Add(e xray.MetricsEvent) {
	key := xray.MetricGroupingKey(e, b.ArgFilter)

	b.Lock()
	defer b.Unlock()

	switch e.GetType() {
	case xray.INCREMENT:
		b.inc.Add(key, e.GetValue())
	case xray.GAUGE:
		b.gauge.Add(key, e.GetValue())
	case xray.DURATION:
		b.delta.Add(key, e.GetValue())
	}

	if _, ok := b.definitions[key]; !ok {
		b.definitions[key] = definition{key: e.GetKey(), args: b.ArgFilter(e.Args())}
	}
}

// FlushTo flushes all data to provided target
func (b *Buffer) FlushTo(target xray.Handler) {
	b.Lock()
	inc := b.inc
	gauge := b.gauge
	delta := b.delta
	definitions := b.definitions

	hasInc := !inc.IsEmpty()
	hasGauge := !gauge.IsEmpty()
	hasDelta := !delta.IsEmpty()

	if len(definitions) > 0 {
		b.definitions = map[string]definition{}

		if hasInc {
			b.inc = NewIncrementalHolder()
		}
		if hasGauge {
			b.gauge = NewGaugeHolder()
		}
		if hasDelta {
			b.delta = NewDeltaHolder(b.percentiles...)
		}
	}
	b.Unlock()

	// Sending inc values
	if hasInc {
		target(inc.Flush(func(groupingKey, _ string, t xray.MetricType, value int64) xray.MetricsEvent {
			return event{
				definition: definitions[groupingKey],
				t:          t,
				value:      value,
			}
		})...)
	}

	// Sending gauge values
	if hasGauge {
		target(gauge.Flush(func(groupingKey, _ string, t xray.MetricType, value int64) xray.MetricsEvent {
			return event{
				definition: definitions[groupingKey],
				t:          t,
				value:      value,
			}
		})...)
	}

	// Sending time delta values
	if hasDelta {
		target(delta.Flush(func(groupingKey, suffix string, t xray.MetricType, value int64) xray.MetricsEvent {
			e := event{
				definition: definitions[groupingKey],
				t:          t,
				value:      value,
			}
			e.key = e.key + suffix
			return e
		})...)
	}
}

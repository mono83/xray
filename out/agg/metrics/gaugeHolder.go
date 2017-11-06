package metrics

import "github.com/mono83/xray"

// NewGaugeHolder builds and returns grouping buffer for gauge values
func NewGaugeHolder() GroupingBuffer {
	return gaugeHolder(map[string]int64{})
}

// gaugeHolder contains grouping buffer for gauge values
type gaugeHolder map[string]int64

func (g gaugeHolder) IsEmpty() bool {
	return len(g) == 0
}

func (g gaugeHolder) Add(groupingKey string, value int64) {
	g[groupingKey] = value
}

func (g gaugeHolder) Flush(build func(string, string, xray.MetricType, int64) xray.MetricsEvent) []xray.Event {
	l := len(g)
	if l == 0 {
		return nil
	}

	events := make([]xray.Event, l)
	j := 0
	for k, v := range g {
		events[j] = build(k, "", xray.GAUGE, v)
		j++
	}

	return events
}

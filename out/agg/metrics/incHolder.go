package metrics

import "github.com/mono83/xray"

// NewIncrementalHolder builds and returns grouping buffer for incremental values
func NewIncrementalHolder() GroupingBuffer {
	return incHolder(map[string]int64{})
}

// incHolder contains grouping buffer for incremental values
type incHolder map[string]int64

func (i incHolder) Add(groupingKey string, value int64) {
	if previous, ok := i[groupingKey]; ok {
		i[groupingKey] = previous + value
	} else {
		i[groupingKey] = value
	}
}

func (i incHolder) IsEmpty() bool {
	return len(i) == 0
}

func (i incHolder) Flush(build func(string, string, xray.MetricType, int64) xray.MetricsEvent) []xray.Event {
	l := len(i)
	if l == 0 {
		return nil
	}

	events := make([]xray.Event, l)
	j := 0
	for k, v := range i {
		events[j] = build(k, "", xray.INCREMENT, v)
		j++
	}

	return events
}

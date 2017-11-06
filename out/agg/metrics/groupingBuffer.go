package metrics

import "github.com/mono83/xray"

// GroupingBuffer describes structures, that can perform data aggregation around groping key
type GroupingBuffer interface {
	// IsEmpty returns true if buffer does not contain any data
	IsEmpty() bool
	// Add method adds new record to grouping buffer
	// This method is not supposed to be thread-safe, so synchronization must be performed externally
	Add(groupingKey string, value int64)
	// Flush return (but not clears) all internal data as metric events using builder func
	// This method is not supposed to be thread-safe, so synchronization must be performed externally
	Flush(func(groupingKey, suffix string, t xray.MetricType, value int64) xray.MetricsEvent) []xray.Event
}

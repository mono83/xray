package prometheus

import (
	"sync"
	"time"

	"github.com/mono83/xray"
)

const inf = 9223372036854775807

// NewExporter instantiates new exporter for Prometheus
func NewExporter(
	f xray.ArgFilter,
	defaultArgs xray.Bucket,
	buckets ...time.Duration,
) *Exporter {
	// Avoiding nil arguments filter
	if f == nil {
		f = func(in []xray.Arg) []xray.Arg {
			return in
		}
	}

	// Avoiding null default args
	if defaultArgs == nil {
		defaultArgs = xray.CreateBucket()
	}

	// Mapping durations into nanos
	nanos := make([]int64, len(buckets)+1)
	for i, j := range buckets {
		nanos[i] = j.Nanoseconds()
	}
	nanos[len(nanos)-1] = inf

	return &Exporter{
		counters:    map[string]*value{},
		gauges:      map[string]*value{},
		histogram:   map[string]*timeValue{},
		buckets:     nanos,
		defaultArgs: defaultArgs,
		filter:      f,
	}
}

// Exporter is Prometheus data exporter
type Exporter struct {
	counters  map[string]*value
	gauges    map[string]*value
	histogram map[string]*timeValue

	buckets     []int64
	defaultArgs xray.Bucket // TODO configure this

	filter xray.ArgFilter
	mutex  sync.Mutex

	cntHandled int64
	cntWritten int64
	cntHTTP    int64
}

// value is an internal structure that holds counters and gaudges data
type value struct {
	value int64
	name  string
	args  []xray.Arg
}

// timeValue is an internal structure that holds time durations
type timeValue struct {
	buckets []int64
	counts  int64
	sum     int64
	name    string
	args    []xray.Arg
}

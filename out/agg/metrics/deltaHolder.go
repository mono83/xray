package metrics

import (
	"sort"

	"github.com/mono83/xray"
	"github.com/mono83/xray/std"
	"strconv"
)

// NewDeltaHolder builds and returns grouping buffer for duration values
// On flushing, this buffer calculates percentiles and other aggregation data
// that is provided instead raw values
func NewDeltaHolder(percentiles ...int) GroupingBuffer {
	if len(percentiles) > 0 {
		for k, v := range percentiles {
			if v < 1 {
				percentiles[k] = 1
			} else if v > 99 {
				percentiles[k] = 99
			}
		}
	}
	return deltaHolder{percentiles: percentiles, values: map[string][]int64{}}
}

type deltaHolder struct {
	percentiles []int
	values      map[string][]int64
}

func (d deltaHolder) Add(groupingKey string, value int64) {
	d.values[groupingKey] = append(d.values[groupingKey], value)
}

func (d deltaHolder) IsEmpty() bool {
	return len(d.values) == 0
}

func (d deltaHolder) Flush(build func(string, string, xray.MetricType, int64) xray.MetricsEvent) []xray.Event {
	l := len(d.values)
	if l == 0 {
		return nil
	}

	events := []xray.Event{}
	for key, values := range d.values {
		var sum, min, max int64
		count := int64(len(values))
		for i, value := range values {
			if i == 0 || value > max {
				max = value
			}
			if i == 0 || value < min {
				min = value
			}
			sum += value
		}

		events = append(
			events,
			build(key, ".count", xray.INCREMENT, count),
			build(key, ".min", xray.GAUGE, min),
			build(key, ".max", xray.GAUGE, max),
			build(key, ".sum", xray.GAUGE, sum),
			build(key, ".avg", xray.GAUGE, sum/count),
		)

		if len(d.percentiles) > 0 {
			sort.Sort(std.Int64Sorter(values))
			for _, percentile := range d.percentiles {
				i := len(values) * percentile / 100
				events = append(events, build(key, ".perc_"+strconv.Itoa(percentile), xray.GAUGE, values[i]))
				sum := int64(0)
				for j := 0; j <= i; j++ {
					sum += values[j]
				}
				events = append(events, build(key, ".mean_"+strconv.Itoa(percentile), xray.GAUGE, sum/int64(i+1)))
			}
		}
	}
	return events
}

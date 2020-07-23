package metrics

import (
	"strconv"

	"github.com/mono83/xray"
	"github.com/mono83/xray/std"
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
		slice := std.Int64Slice(values)
		count, min, max, avg, sum := slice.Analyze()

		events = append(
			events,
			build(key, ".count", xray.INCREMENT, int64(count)),
			build(key, ".min", xray.GAUGE, min),
			build(key, ".max", xray.GAUGE, max),
			build(key, ".sum", xray.GAUGE, sum),
			build(key, ".avg", xray.GAUGE, avg),
		)

		if len(d.percentiles) > 0 {
			slice.Sort()
			for _, percentile := range d.percentiles {
				value, sub := slice.Percentile(percentile)
				sum := std.AggregateSum(sub)

				events = append(events, build(key, ".perc_"+strconv.Itoa(percentile), xray.GAUGE, value))
				if len(sub) > 0 {
					events = append(events, build(key, ".mean_"+strconv.Itoa(percentile), xray.GAUGE, sum/int64(len(sub))))
				} else {
					events = append(events, build(key, ".mean_"+strconv.Itoa(percentile), xray.GAUGE, 0))
				}
			}
		}
	}
	return events
}

package metrics

import (
	"fmt"
	"github.com/mono83/xray"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestDeltaHolder(t *testing.T) {
	assert := assert2.New(t)
	d := NewDeltaHolder(0, 50, 80, 200)

	d.Add("foo", 10)
	d.Add("foo", 20)
	d.Add("foo", 30)
	d.Add("foo", 40)
	d.Add("foo", 33)
	d.Add("foo", 12)
	d.Add("foo", 8)
	d.Add("foo", 31)
	d.Add("foo", 15)
	d.Add("foo", 1)

	events := d.Flush(func(i, s string, metricType xray.MetricType, i2 int64) xray.MetricsEvent {
		return event{definition: definition{key: i + s}, t: metricType, value: i2}
	})

	fmt.Println(events)
	if assert.Len(events, 13) {
		assertEvent(assert, "foo.count", 10, events[0])
		assertEvent(assert, "foo.min", 1, events[1])
		assertEvent(assert, "foo.max", 40, events[2])
		assertEvent(assert, "foo.sum", 200, events[3])
		assertEvent(assert, "foo.avg", 20, events[4])

		assertEvent(assert, "foo.perc_1", 1, events[5])
		assertEvent(assert, "foo.mean_1", 1, events[6])

		assertEvent(assert, "foo.perc_50", 20, events[7])
		assertEvent(assert, "foo.mean_50", 11, events[8])

		assertEvent(assert, "foo.perc_80", 33, events[9])
		assertEvent(assert, "foo.mean_80", 17, events[10])

		assertEvent(assert, "foo.perc_99", 40, events[11])
		assertEvent(assert, "foo.mean_99", 20, events[12])
	}
}

func assertEvent(assert *assert2.Assertions, expectedKey string, expectedValue int64, e xray.Event) {
	if m, ok := e.(xray.MetricsEvent); assert.True(ok) {
		assert.Equal(expectedKey, m.GetKey())
		assert.Equal(expectedValue, m.GetValue())
	}
}

func BenchmarkDelta_NoPercentiles_1000Items(b *testing.B) {
	d := NewDeltaHolder()

	for i := 0; i < 1000; i++ {
		d.Add("foo", int64(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Flush(func(string, string, xray.MetricType, int64) xray.MetricsEvent {
			return event{}
		})
	}
}

func BenchmarkDelta_10Percentiles_1000Items(b *testing.B) {
	d := NewDeltaHolder(0, 10, 20, 30, 40, 50, 60, 70, 80, 90)

	for i := 0; i < 1000; i++ {
		d.Add("foo", int64(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Flush(func(string, string, xray.MetricType, int64) xray.MetricsEvent {
			return event{}
		})
	}
}

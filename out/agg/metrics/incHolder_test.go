package metrics

import (
	"testing"

	"github.com/mono83/xray"
	assert2 "github.com/stretchr/testify/assert"
)

func TestIncHolder(t *testing.T) {
	assert := assert2.New(t)

	buf := NewIncrementalHolder()
	buf.Add("foo", 1)
	buf.Add("foo", 2)
	buf.Add("foo", 3)
	buf.Add("bar", 4)
	buf.Add("bar", 5)

	events := buf.Flush(func(groupingKey, suffix string, t xray.MetricType, value int64) xray.MetricsEvent {
		return event{
			definition: definition{key: groupingKey},
			t:          t,
			value:      value,
		}
	})

	assert.Len(events, 2)
	e1, _ := events[0].(xray.MetricsEvent)
	e2, _ := events[1].(xray.MetricsEvent)
	assert.Equal("foo", e1.GetKey())
	assert.Equal(int64(6), e1.GetValue())
	assert.Equal("bar", e2.GetKey())
	assert.Equal(int64(9), e2.GetValue())
}

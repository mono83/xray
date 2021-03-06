package args

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var durationDataProvider = []struct {
	expected string
	given    time.Duration
}{
	{"0.001s", time.Millisecond},
	{"0.012s", time.Millisecond * 12},
	{"0.810s", time.Millisecond * 810},
	{"5.678s", time.Millisecond * 5678},
	{"56.789s", time.Millisecond * 56789},
	{"9m27s", time.Millisecond * 567890},
	{"15h46m29s", time.Millisecond * 56789000},
}

func TestDurationFormatting(t *testing.T) {
	for _, data := range durationDataProvider {
		t.Run(data.expected, func(t *testing.T) {
			assert.Equal(t, data.expected, Delta(data.given).Value())
			assert.Equal(t, data.expected, Elapsed(data.given).Value())
		})
	}
}

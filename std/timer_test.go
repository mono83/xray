package std

import (
	"testing"
	"time"

	"github.com/mono83/xray/args"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	// Building timer object without ray
	o := Timer(nil, "foo", args.Name("test"))
	assert.NotNil(t, o)

	// Cast
	if x, ok := o.(*timer); assert.True(t, ok) {
		// Initial state
		assert.False(t, x.start.IsZero())
		assert.True(t, x.stop.IsZero())

		// Stopping
		_ = <-time.After(time.Microsecond)
		o.Stop()
		assert.False(t, x.start.IsZero())
		assert.False(t, x.stop.IsZero())

		current := x.stop

		// Double stop
		o.Stop()
		assert.Equal(t, current.UnixNano(), x.stop.UnixNano())
		assert.True(t, o.Stop().Nanoseconds() >= 1000)
	}
}

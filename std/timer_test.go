package std

import (
	"github.com/mono83/xray/args"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestTimer(t *testing.T) {
	assert := assert2.New(t)

	// Building timer object without ray
	o := Timer(nil, "foo", args.Name("test"))
	assert.NotNil(o)

	// Cast
	if x, ok := o.(*timer); ok {
		// Initial state
		assert.False(x.start.IsZero())
		assert.True(x.stop.IsZero())

		// Stopping
		o.Stop()
		assert.False(x.start.IsZero())
		assert.False(x.stop.IsZero())

		current := x.stop

		// Double stop
		o.Stop()
		assert.Equal(current.UnixNano(), x.stop.UnixNano())
		assert.True(o.Stop().Nanoseconds() > 0)
	} else {
		assert.Fail("Not a timer struct")
	}
}

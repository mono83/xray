package std

import (
	"time"

	"github.com/mono83/xray"
)

// TimeHelper is helper structure, used to measure elapsed time with automatic data forwarding
// into associated ray
type TimeHelper interface {
	xray.NanoHolder
	Stop() time.Duration
}

// Timer builds new timer (TimeHelper instance) used for time measurement
func Timer(r xray.Ray, name string, args ...xray.Arg) TimeHelper {
	return &timer{
		start: time.Now(),
		name:  name,
		ray:   r,
		args:  args,
	}
}

type timer struct {
	start, stop time.Time
	name        string
	ray         xray.Ray
	args        []xray.Arg
}

func (t *timer) Stop() time.Duration {
	if t.stop.IsZero() {
		t.stop = time.Now()
		if t.ray != nil {
			t.ray.Duration(t.name, t.stop.Sub(t.start), t.args...)
		}
	}

	return t.stop.Sub(t.start)
}

func (t *timer) Nanoseconds() int64 {
	return t.Stop().Nanoseconds()
}

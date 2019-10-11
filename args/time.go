package args

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	sec := d.Seconds()
	if sec >= 60 {
		// Truncating duration to seconds
		return time.Duration((d / time.Second) * time.Second).String()
	}
	return fmt.Sprintf("%.3fs", sec)
}

// Delta arg contains duration of some event with name "delta"
type Delta time.Duration

// Name is ray.Arg interface implementation. Returns argument name
func (d Delta) Name() string { return "delta" }

// Value is ray.Arg interface implementation. Returns argument value
func (d Delta) Value() string { return formatDuration(d.Duration()) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (d Delta) Scalar() interface{} { return d.Duration().Nanoseconds() }

// Duration returns time.Duration stored within arg
func (d Delta) Duration() time.Duration { return time.Duration(d) }

// Elapsed arg contains duration of some event with name "elapsed"
type Elapsed time.Duration

// Name is ray.Arg interface implementation. Returns argument name
func (e Elapsed) Name() string { return "elapsed" }

// Value is ray.Arg interface implementation. Returns argument value
func (e Elapsed) Value() string { return formatDuration(e.Duration()) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (e Elapsed) Scalar() interface{} { return e.Duration().Nanoseconds() }

// Duration returns time.Duration stored within arg
func (e Elapsed) Duration() time.Duration { return time.Duration(e) }

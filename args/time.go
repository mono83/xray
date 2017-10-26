package args

import (
	"fmt"
	"time"
)

// Delta arg contains duration of some event with name "delta"
type Delta time.Duration

// Name is ray.Arg interface implementation. Returns argument name
func (d Delta) Name() string { return "delta" }

// Value is ray.Arg interface implementation. Returns argument value
func (d Delta) Value() string { return fmt.Sprintf("%.3fs", time.Duration(d).Seconds()) }

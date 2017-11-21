package env

import (
	"os"
	"strconv"
)

// PID contains current process identifier
var PID ArgPID

// ArgPID is process identifier, also compatible with xray.Arg
type ArgPID int

// Int returns integer representation of pid
func (a ArgPID) Int() int { return int(a) }

// Name returns argument key (name)
func (a ArgPID) Name() string { return "pid" }

// Value returns string representation of argument value
func (a ArgPID) Value() string { return strconv.Itoa(a.Int()) }

// Scalar returns raw representation of argument value. It can be scalar value or slice of scalar values.
func (a ArgPID) Scalar() interface{} { return a.Int() }

func init() {
	PID = ArgPID(os.Getpid())
}

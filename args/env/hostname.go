package env

import "os"

// HostName contains host name of current machine
var HostName ArgHostName

// ArgHostName is host name, also compatible with xray.Arg
type ArgHostName string

// Int returns integer representation of pid
func (a ArgHostName) String() string { return string(a) }

// Name returns argument key (name)
func (ArgHostName) Name() string { return "hostname" }

// Value returns string representation of argument value
func (a ArgHostName) Value() string { return string(a) }

// Scalar returns raw representation of argument value. It can be scalar value or slice of scalar values.
func (a ArgHostName) Scalar() interface{} { return string(a) }

func init() {
	if name, err := os.Hostname(); err == nil {
		HostName = ArgHostName(name)
	} else {
		HostName = ArgHostName("unknown")
	}
}

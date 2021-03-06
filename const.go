package xray

import (
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/id"
)

// MetricType describes metric type
type MetricType byte

// Level describes logging level
type Level byte

// DumpSource describes dump source
type DumpSource byte

// List of defined log levels
const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	ALERT
	CRITICAL
)

// List of defined metric types
const (
	INCREMENT MetricType = iota
	GAUGE
	DURATION
)

// List of defined dump sources
const (
	IN DumpSource = iota
	OUT
)

// ROOT is main top-level ray. Can be used to attach listeners
var ROOT Ray

// BOOT is ray used to log boot operations
var BOOT Ray

// OrRoot checks if candidate is nil.
// Returns candidate itself or ROOT ray if candidate is nil.
func OrRoot(candidate Ray) Ray {
	if candidate == nil {
		return ROOT
	}

	return candidate
}

// OrRootFork checks if candidate is nil.
// Returns candidate itself or ROOT ray fork if candidate is nil.
func OrRootFork(candidate Ray) Ray {
	if candidate == nil {
		return ROOT.Fork()
	}

	return candidate
}

func init() {
	ROOT = New(NewSyncEmitter, id.Generator20Base64).WithLogger("ROOT").With(args.Host(id.Hostname))
	BOOT = ROOT.Fork().WithLogger("BOOT")
}

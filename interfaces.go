package xray

import "time"

// Arg describes ray logging qualifier (argument)
type Arg interface {
	Name() string
	Value() string
}

// Bucket is container for arguments
type Bucket interface {
	// Size return number of args within bucket
	Size() int
	// Get returns argument (or nil) by key
	Get(string) Arg
	// Args returns args slice
	Args() []Arg
}

// Handler is events handler function
type Handler func(...Event)

// EventEmitter describes common event emitter interface
type EventEmitter interface {
	// Emit method emits event into ray
	Emit(...Event)

	// On registers listeners of logging events
	On(Handler)
}

// ExtendedEmitter is extended event emitter, that has helper methods to build events
type ExtendedEmitter interface {
	EventEmitter

	Inc(string, ...Arg)
	Increment(string, int64, ...Arg)
	Gauge(string, int64, ...Arg)
	Duration(string, time.Duration, ...Arg)

	Pass(error) error

	Trace(string, ...Arg)
	Debug(string, ...Arg)
	Info(string, ...Arg)
	Warning(string, ...Arg)
	Error(string, ...Arg)
	Alert(string, ...Arg)
	Critical(string, ...Arg)
}

// RayIDProvider describes components that are able to return (or even generate) ray IDs
type RayIDProvider interface {
	GetRayID() string
}

// Ray describes execution context (ray)
type Ray interface {
	RayIDProvider
	ExtendedEmitter

	// GetLogger returns configured logger name
	GetLogger() string

	// GetMetricPrefix returns configured metrics prefix
	GetMetricPrefix() string

	// GetArguments returns full container with arguments
	GetArguments() Bucket

	// WithLogger returns clone of ray interface with some name applied
	WithLogger(string) Ray

	// WithMetricPrefix returns clone of ray interface with some name applied
	WithMetricPrefix(string) Ray

	// With method rebuilds ray with provided arguments
	With(...Arg) Ray

	// Creates new ray with new unique ID
	Fork() Ray
}

// Event describes ray logging event
type Event interface {
	RayIDProvider
	Bucket
}

// LogEvent represents logging event
type LogEvent interface {
	Event

	// GetTime returns event generation time
	GetTime() time.Time

	// GetLevel returns logging level
	GetLevel() Level

	// GetLogger returns logger name
	GetLogger() string

	// GetMessage returns string to log
	GetMessage() string
}

// MetricsEvent represents metrics event
type MetricsEvent interface {
	Event

	// GetType returns metric type
	GetType() MetricType
	// GetKey returns metrics key
	GetKey() string
	// GetValue return metrics value
	GetValue() int64
}

// ByteDumpEvent used to provide dumps
type ByteDumpEvent interface {
	Event

	GetSource() DumpSource
	GetData() []byte
}

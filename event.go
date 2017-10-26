package xray

// MetricType describes metric type
type MetricType byte

// Level describes logging level
type Level byte

// List of defined log levels
const (
	TRACE    Level = iota
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

// Event describes ray logging event
type Event interface {
	RayIDProvider
	Bucket
}

// LogEvent represents logging event
type LogEvent interface {
	Event

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

	// GetKey returns metrics key
	GetKey() string
	// GetValue return metrics value
	GetValue() int64
}

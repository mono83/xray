package ray

// Event describes ray logging event
type Event interface {
	// Splits separates event into logging and metrics ones.
	// Both can be nil and both can be set
	Split() ([]LogEvent, []MetricsEvent)

	// Args returns map of arguments for event
	Args() []Arg
}

// LogEvent represents logging event
type LogEvent interface {
	Event

	// Message returns string to log
	Message() string
}

// MetricsEvent represents metrics event
type MetricsEvent interface {
	Event

	Key() string
	Value() int64
}

// EmittedEvent describes emitted event
type EmittedEvent struct {
	Event
	Emitter Interface
	Depth   int
}

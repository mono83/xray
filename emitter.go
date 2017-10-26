package xray

// EventEmitter describes common event emitter interface
type EventEmitter interface {
	// Emit method emits event into ray
	Emit(...Event)

	// On registers listeners of logging events
	On(func(...Event))
}

// ExtendedEmitter is extended event emitter, that has helper methods to build events
type ExtendedEmitter interface {
	EventEmitter

	Trace(string, ...Arg)
	Debug(string, ...Arg)
	Info(string, ...Arg)
	Warning(string, ...Arg)
	Error(string, ...Arg)
	Alert(string, ...Arg)
	Critical(string, ...Arg)
}

package ray

// Interface describes execution context (ray)
type Interface interface {
	// GetID returns unique identifier of ray
	GetID() string

	// Args returns slice of arguments for event
	Args() []Arg

	// Group returns clone of ray interface with some name applied
	Group(string) Interface

	// Emit method emits event into ray
	Emit(...Event)

	// On registers listeners of logging events
	On(func(EmittedEvent))

	// Creates new ray with new unique ID
	Fork() Interface
}

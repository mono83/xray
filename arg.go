package ray

// Arg describes ray logging qualifier (argument)
type Arg interface {
	Name() string
	Value() string
}

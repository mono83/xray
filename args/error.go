package args

// Error is special arg, used to hold errors with key "err"
type Error struct {
	Err error
}

// Name is xray.Arg interface implementation. Returns argument name
func (Error) Name() string { return "err" }

// Value is xray.Arg interface implementation. Returns argument value
func (e Error) Value() string { return e.Err.Error() }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (e Error) Scalar() interface{} { return e.Err.Error() }

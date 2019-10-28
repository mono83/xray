package args

// Nil is special case arg, that represents nil values
type Nil string

// Name is xray.Arg interface implementation. Returns argument name
func (n Nil) Name() string { return string(n) }

// Value is xray.Arg interface implementation. Returns argument value
func (Nil) Value() string { return "" }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (Nil) Scalar() interface{} { return nil }

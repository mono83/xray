package args

import "strconv"

// Int64 is common arg, that contain arbitrary int64
type Int64 struct {
	N string
	V int64
}

// Name is ray.Arg interface implementation. Returns argument name
func (i Int64) Name() string { return i.N }

// Value is ray.Arg interface implementation. Returns argument value
func (i Int64) Value() string { return strconv.FormatInt(i.V, 10) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (i Int64) Scalar() interface{} { return i.V }

// Count represents integer arg with name "count"
type Count int

// Name is ray.Arg interface implementation. Returns argument name
func (c Count) Name() string { return "count" }

// Value is ray.Arg interface implementation. Returns argument value
func (c Count) Value() string { return strconv.Itoa(int(c)) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (c Count) Scalar() interface{} { return int(c) }

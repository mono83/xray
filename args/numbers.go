package args

import "strconv"

// Int64 is common arg, that contain arbitrary int64
type Int64 struct {
	N string
	V int64
}

// Name is xray.Arg interface implementation. Returns argument name
func (i Int64) Name() string { return i.N }

// Value is xray.Arg interface implementation. Returns argument value
func (i Int64) Value() string { return strconv.FormatInt(i.V, 10) }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (i Int64) Scalar() interface{} { return i.V }

// ID64 is number argument with name "id"
type ID64 int64

// Name is xray.Arg interface implementation. Returns argument name
func (ID64) Name() string { return "id" }

// Value is xray.Arg interface implementation. Returns argument value
func (i ID64) Value() string { return strconv.FormatInt(int64(i), 10) }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (i ID64) Scalar() interface{} { return int64(i) }

// Int is common arg, that contain arbitrary int value
type Int struct {
	N string
	V int
}

// Name is xray.Arg interface implementation. Returns argument name
func (i Int) Name() string { return i.N }

// Value is xray.Arg interface implementation. Returns argument value
func (i Int) Value() string { return strconv.Itoa(i.V) }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (i Int) Scalar() interface{} { return i.V }

// Count represents integer arg with name "count"
type Count int

// Name is xray.Arg interface implementation. Returns argument name
func (c Count) Name() string { return "count" }

// Value is xray.Arg interface implementation. Returns argument value
func (c Count) Value() string { return strconv.Itoa(int(c)) }

// Scalar is xray.Arg interface implementation. Returns argument value as scalar
func (c Count) Scalar() interface{} { return int(c) }

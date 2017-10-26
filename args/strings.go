package args

// String is common arg, that contain arbitrary string
type String struct {
	N, V string
}

// Name is ray.Arg interface implementation. Returns argument name
func (s String) Name() string { return s.N }

// Value is ray.Arg interface implementation. Returns argument value
func (s String) Value() string { return s.V }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (s String) Scalar() interface{} { return s.V }

// SQL is string argument with name "sql"
type SQL string

// Name is ray.Arg interface implementation. Returns argument name
func (s SQL) Name() string { return "sql" }

// Value is ray.Arg interface implementation. Returns argument value
func (s SQL) Value() string { return string(s) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (s SQL) Scalar() interface{} { return string(s) }

// Name is string argument with name "name"
type Name string

// Name is ray.Arg interface implementation. Returns argument name
func (n Name) Name() string { return "name" }

// Value is ray.Arg interface implementation. Returns argument value
func (n Name) Value() string { return string(n) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (n Name) Scalar() interface{} { return string(n) }

// RayID is string argument with name "rayId"
type RayID string

// Name is ray.Arg interface implementation. Returns argument name
func (r RayID) Name() string { return "rayId" }

// Value is ray.Arg interface implementation. Returns argument value
func (r RayID) Value() string { return string(r) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (r RayID) Scalar() interface{} { return string(r) }

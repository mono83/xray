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

// Type is string argument with name "type"
type Type string

// Name is ray.Arg interface implementation. Returns argument name
func (t Type) Name() string { return "type" }

// Value is ray.Arg interface implementation. Returns argument value
func (t Type) Value() string { return string(t) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (t Type) Scalar() interface{} { return string(t) }

// Name is string argument with name "name"
type Name string

// Name is ray.Arg interface implementation. Returns argument name
func (n Name) Name() string { return "name" }

// Value is ray.Arg interface implementation. Returns argument value
func (n Name) Value() string { return string(n) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (n Name) Scalar() interface{} { return string(n) }

// AppName is string argument with name "app"
type AppName string

// Name is ray.Arg interface implementation. Returns argument name
func (a AppName) Name() string { return "app" }

// Value is ray.Arg interface implementation. Returns argument value
func (a AppName) Value() string { return string(a) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (a AppName) Scalar() interface{} { return string(a) }

// RayID is string argument with name "rayId"
type RayID string

// Name is ray.Arg interface implementation. Returns argument name
func (r RayID) Name() string { return "rayId" }

// Value is ray.Arg interface implementation. Returns argument value
func (r RayID) Value() string { return string(r) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (r RayID) Scalar() interface{} { return string(r) }

// Host is string argument with name "host"
type Host string

// Name is ray.Arg interface implementation. Returns argument name
func (h Host) Name() string { return "host" }

// Value is ray.Arg interface implementation. Returns argument value
func (h Host) Value() string { return string(h) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (h Host) Scalar() interface{} { return string(h) }

// URL is string argument with name "url"
type URL string

// Name is ray.Arg interface implementation. Returns argument name
func (u URL) Name() string { return "url" }

// Value is ray.Arg interface implementation. Returns argument value
func (u URL) Value() string { return string(u) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (u URL) Scalar() interface{} { return string(u) }

// Addr is string argument with name "addr"
type Addr string

// Name is ray.Arg interface implementation. Returns argument name
func (a Addr) Name() string { return "addr" }

// Value is ray.Arg interface implementation. Returns argument value
func (a Addr) Value() string { return string(a) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (a Addr) Scalar() interface{} { return string(a) }

// Method is string argument with name "method"
type Method string

// Name is ray.Arg interface implementation. Returns argument name
func (m Method) Name() string { return "method" }

// Value is ray.Arg interface implementation. Returns argument value
func (m Method) Value() string { return string(m) }

// Scalar is ray.Arg interface implementation. Returns argument value as scalar
func (m Method) Scalar() interface{} { return string(m) }

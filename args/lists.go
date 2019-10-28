package args

import (
	"strconv"
	"strings"
)

// ID64List contains multiple IDs
type ID64List []int64

// Name is xray.Arg interface implementation. Returns argument name
func (i ID64List) Name() string { return "id" }

// Value is xray.Arg interface implementation. Returns argument value
func (i ID64List) Value() string {
	return strings.Join(i.ValueList(), ",")
}

// ValueList returns values as string slice
func (i ID64List) ValueList() []string {
	l := len(i)
	if l == 0 {
		return []string{}
	} else if l == 1 {
		return []string{strconv.FormatInt(i[0], 10)}
	}

	str := make([]string, l)
	for k, v := range i {
		str[k] = strconv.FormatInt(v, 10)
	}

	return str
}

// Scalar is xray.Arg interface implementation. Returns argument value as slice
func (i ID64List) Scalar() interface{} { return []int64(i) }

// NameList contains multiple names
type NameList []string

// Name is xray.Arg interface implementation. Returns argument name
func (n NameList) Name() string { return "name" }

// Value is xray.Arg interface implementation. Returns argument value
func (n NameList) Value() string {
	return strings.Join(n, ",")
}

// ValueList returns values as string slice
func (n NameList) ValueList() []string {
	return []string(n)
}

// Scalar is xray.Arg interface implementation. Returns argument value as slice
func (n NameList) Scalar() interface{} { return []string(n) }

// TypeList contains multiple types
type TypeList []string

// Name is xray.Arg interface implementation. Returns argument name
func (t TypeList) Name() string { return "type" }

// Value is xray.Arg interface implementation. Returns argument value
func (t TypeList) Value() string {
	return strings.Join(t, ",")
}

// ValueList returns values as string slice
func (t TypeList) ValueList() []string {
	return []string(t)
}

// Scalar is xray.Arg interface implementation. Returns argument value as slice
func (t TypeList) Scalar() interface{} { return []string(t) }

package xray

import "context"

// contextKeyType
type contextKeyType int8

// contextKey contains key to match xray.Ray from Golang context
const contextKey contextKeyType = 1

// contextWithRay contains Golang context with xray.Ray embedded
type contextWithRay struct {
	context.Context
	ray Ray
}

// Value is context.Context interface implementation
func (c contextWithRay) Value(key interface{}) interface{} {
	if key == contextKey {
		return c.ray
	}

	return c.Context.Value(key)
}

// WrapContext wraps given Golang context producing new one with xray.Ray embedded
func WrapContext(c context.Context, r Ray) context.Context {
	if cr, ok := c.(contextWithRay); ok {
		// Extracting nested context
		c = cr.Context
	}

	return contextWithRay{
		Context: c,
		ray:     OrRootFork(r),
	}
}

// FromContext extracts xray.Ray from given Golang context.
// Will produce nil if there is no context.
func FromContext(c context.Context) Ray {
	if c != nil {
		v := c.Value(contextKey)
		if r, ok := v.(Ray); ok {
			return r
		}
	}

	return nil
}

// FromContextOr extracts xray.Ray from given Golang context.
// If there were no xray.Ray in context will return given fallback ray.
// Further more if even given fallback ray is nil, func will fork ROOT, so in
// any case this method will not return nil.
func FromContextOr(c context.Context, fallback Ray) Ray {
	if r := FromContext(c); r != nil {
		return r
	}

	return OrRootFork(fallback)
}

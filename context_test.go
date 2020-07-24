package xray

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapContext(t *testing.T) {
	w := WrapContext(context.TODO(), ROOT)
	_, ok := w.(contextWithRay)
	assert.True(t, ok)

	w2 := WrapContext(w, ROOT)
	cwr, ok := w2.(contextWithRay)
	if assert.True(t, ok) {
		_, ok = cwr.Context.(contextWithRay)
		assert.False(t, ok)
	}
}

func TestContextValue(t *testing.T) {
	w := WrapContext(context.TODO(), BOOT)

	v := w.Value(contextKey)
	if assert.NotNil(t, v) {
		_, ok := v.(Ray)
		assert.True(t, ok)
	}

	// Reading int8 instead of contextKey
	v = w.Value(int8(1))
	assert.Nil(t, v)
}

func TestFromContext(t *testing.T) {
	r := FromContext(context.TODO())
	assert.Nil(t, r)

	r = FromContext(WrapContext(context.TODO(), ROOT))
	assert.NotNil(t, r)
}

func TestFromContextOr(t *testing.T) {
	r := FromContextOr(context.TODO(), BOOT)
	assert.NotNil(t, r)
	assert.Equal(t, BOOT, r)
	assert.NotEqual(t, ROOT, r)

	r = FromContextOr(context.TODO(), nil)
	assert.NotNil(t, r)
	assert.NotEqual(t, BOOT, r)
	assert.NotEqual(t, ROOT, r)

	r = FromContextOr(WrapContext(context.TODO(), BOOT), nil)
	assert.NotNil(t, r)
	assert.Equal(t, BOOT, r)
	assert.NotEqual(t, ROOT, r)
}

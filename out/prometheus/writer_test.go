package prometheus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEscape(t *testing.T) {
	assert.Equal(t, string(escape("foo-bar.baz")), "foo_bar_baz")
	assert.Equal(t, string(escape("hello world")), "hello_world")
	assert.Equal(t, string(escape(" trim spaces  ")), "trim_spaces")
}

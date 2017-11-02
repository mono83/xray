package xray

import (
	"github.com/mono83/xray/args"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupingKeyArgs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("", GroupingKeyArgs(nil))
	assert.Equal("", GroupingKeyArgs([]Arg{}))
	assert.Equal("name\tfoo", GroupingKeyArgs([]Arg{args.Name("foo")}))
	assert.Equal("foo\t", GroupingKeyArgs([]Arg{args.Nil("foo")}))
	assert.Equal("id\t-1\tname\tfoo\ttype\tbar", GroupingKeyArgs([]Arg{args.Name("foo"), args.Type("bar"), args.ID64(-1)}))
}

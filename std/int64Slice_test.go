package std

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestInt64Sorter(t *testing.T) {
	assert := assert.New(t)

	data := []int64{4, 1, 8, -2}
	sort.Sort(Int64Slice(data))

	assert.Equal(int64(-2), data[0])
	assert.Equal(int64(1), data[1])
	assert.Equal(int64(4), data[2])
	assert.Equal(int64(8), data[3])
}

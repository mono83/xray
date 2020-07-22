package std

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccumulatorInt64(t *testing.T) {
	assert := assert.New(t)

	acc := &AccumulatorInt64{Depth: 2, Aggregate: AggregateSum}

	// Adding
	for i := 1; i <= 14; i++ {
		acc.Add(int64(i))
		assert.Equal(int64(i), acc.List()[0])
	}

	values := acc.List()
	assert.Len(values, 6)
	assert.Equal(int64(14), values[0])
	assert.Equal(int64(13), values[1])
	assert.Equal(int64(12+11), values[2])
	assert.Equal(int64(10+9), values[3])
	assert.Equal(int64(8+7+6+5), values[4])
	assert.Equal(int64(4+3+2+1), values[5])
}

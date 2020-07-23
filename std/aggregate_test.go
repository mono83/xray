package std

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregations(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(0), AggregateSum([]int64{}))
	assert.Equal(int64(0), AggregateAvg([]int64{}))
	assert.Equal(int64(0), AggregateMax([]int64{}))
	assert.Equal(int64(0), AggregateMin([]int64{}))

	assert.Equal(int64(0), AggregateSum(nil))
	assert.Equal(int64(0), AggregateAvg(nil))
	assert.Equal(int64(0), AggregateMax(nil))
	assert.Equal(int64(0), AggregateMin(nil))

	assert.Equal(int64(30), AggregateSum([]int64{10, 5, 15}))
	assert.Equal(int64(10), AggregateAvg([]int64{10, 5, 15}))
	assert.Equal(int64(15), AggregateMax([]int64{10, 5, 15}))
	assert.Equal(int64(5), AggregateMin([]int64{10, 5, 15}))
}

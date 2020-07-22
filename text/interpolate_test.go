package text

import (
	"testing"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/stretchr/testify/assert"
)

func TestInterpolatePlainText(t *testing.T) {
	assert := assert.New(t)

	bucket := xray.CreateBucket(
		args.Name("tester"),
		args.Int64{N: "id", V: 15},
	)

	assert.Equal("This is tester, id 15 ", InterpolatePlainText("This is :name, id :id :foo", bucket, false))
	assert.Equal("This is [tester], id [15] <!foo!>", InterpolatePlainText("This is :name, id :id :foo", bucket, true))
}

func TestInterpolateLists(t *testing.T) {
	assert := assert.New(t)
	bucket := xray.CreateBucket(
		args.NameList([]string{"foo", "bar"}),
		args.TypeList([]string{"alpha", "beta", "gamma"}),
		args.ID64List([]int64{-5, 16, 0}),
	)

	assert.Equal("For foo and bar", InterpolatePlainText("For :name and :name", bucket, false))
	assert.Equal("alpha beta gamma", InterpolatePlainText(":type :type :type", bucket, false))
	assert.Equal("From -5 to 16 using 0", InterpolatePlainText("From :id to :id using :id", bucket, false))
}

func TestInterpolateListsOne(t *testing.T) {
	assert := assert.New(t)
	bucket := xray.CreateBucket(
		args.NameList([]string{"foo"}),
		args.TypeList([]string{"alpha"}),
		args.ID64List([]int64{-5}),
	)

	assert.Equal("For foo and ", InterpolatePlainText("For :name and :name", bucket, false))
	assert.Equal("alpha  ", InterpolatePlainText(":type :type :type", bucket, false))
	assert.Equal("From -5 to  using ", InterpolatePlainText("From :id to :id using :id", bucket, false))
}

func TestInterpolateListsEmpty(t *testing.T) {
	assert := assert.New(t)
	bucket := xray.CreateBucket(
		args.NameList(nil),
		args.TypeList([]string{}),
	)

	assert.Equal("For  and ", InterpolatePlainText("For :name and :name", bucket, false))
	assert.Equal("  ", InterpolatePlainText(":type :type :type", bucket, false))
}

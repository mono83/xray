package text

import (
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/stretchr/testify/assert"
	"testing"
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

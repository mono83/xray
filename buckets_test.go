package xray

import (
	"github.com/mono83/xray/args"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBucket(t *testing.T) {
	assert := assert.New(t)

	var bucket Bucket = emptyBucket{}
	assert.Equal(0, bucket.Size())
	assert.Nil(bucket.Args())
	assert.Nil(bucket.Get("foo"))
}

func TestSingleArgBucket(t *testing.T) {
	assert := assert.New(t)

	var bucket Bucket = singleArgBucket{args.String{N: "foo", V: "bar"}}
	assert.Equal(1, bucket.Size())
	assert.Len(bucket.Args(), 1)
	assert.Equal("bar", bucket.Get("foo").Value())
	assert.Nil(bucket.Get("bar"))
}

func TestMapBucket(t *testing.T) {
	assert := assert.New(t)

	m := map[string]Arg{}
	m["foo"] = args.String{N: "foo", V: "bar"}
	m["baz"] = args.Int64{N: "baz", V: 100500}

	var bucket Bucket = mapBucket(m)

	assert.Equal(2, bucket.Size())
	assert.Len(bucket.Args(), 2)
	assert.Equal("bar", bucket.Get("foo").Value())
	assert.Equal("100500", bucket.Get("baz").Value())
	assert.Nil(bucket.Get("bar"))
}

func TestCreateBucket(t *testing.T) {
	assert := assert.New(t)

	bucket := CreateBucket()
	assert.IsType(emptyBucket{}, bucket)

	bucket = CreateBucket(nil)
	assert.IsType(emptyBucket{}, bucket)

	bucket = CreateBucket(args.String{N: "foo", V: "bar"})
	assert.IsType(singleArgBucket{}, bucket)

	bucket = CreateBucket(args.String{N: "foo", V: ""}, args.String{N: "bar", V: ""})
	assert.IsType(mapBucket{}, bucket)
}

func TestAppendBucket(t *testing.T) {
	assert := assert.New(t)

	bucket := CreateBucket()
	assert.IsType(emptyBucket{}, AppendBucket(bucket))
	assert.IsType(singleArgBucket{}, AppendBucket(bucket, args.String{N: "foo", V: ""}))
	assert.IsType(mapBucket{}, AppendBucket(bucket, args.String{N: "foo", V: ""}, args.String{N: "bar", V: ""}))

	bucket = CreateBucket(args.String{N: "foo", V: ""})
	assert.IsType(singleArgBucket{}, AppendBucket(bucket))
	assert.IsType(mapBucket{}, AppendBucket(bucket, args.String{N: "foo", V: ""}))
	assert.IsType(mapBucket{}, AppendBucket(bucket, args.String{N: "foo", V: ""}, args.String{N: "bar", V: ""}))
}

func BenchmarkCreateBucketEmpty(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		CreateBucket()
	}
}

func BenchmarkCreateBucketSingle(b *testing.B) {
	arg := args.String{N: "foo", V: ""}
	for i := 0; i <= b.N; i++ {
		CreateBucket(arg)
	}
}

func BenchmarkCreateBucketMap2(b *testing.B) {
	aa := []Arg{
		args.String{N: "foo1", V: ""},
		args.String{N: "foo2", V: ""},
	}
	for i := 0; i <= b.N; i++ {
		CreateBucket(aa...)
	}
}

func BenchmarkCreateBucketMap15(b *testing.B) {
	aa := []Arg{
		args.String{N: "foo1", V: ""},
		args.String{N: "foo2", V: ""},
		args.String{N: "foo3", V: ""},
		args.String{N: "foo4", V: ""},
		args.String{N: "foo5", V: ""},
		args.String{N: "foo6", V: ""},
		args.String{N: "foo7", V: ""},
		args.String{N: "foo8", V: ""},
		args.String{N: "foo9", V: ""},
		args.String{N: "foo10", V: ""},
		args.String{N: "foo11", V: ""},
		args.String{N: "foo12", V: ""},
		args.String{N: "foo13", V: ""},
		args.String{N: "foo14", V: ""},
		args.String{N: "foo15", V: ""},
	}
	for i := 0; i <= b.N; i++ {
		CreateBucket(aa...)
	}
}

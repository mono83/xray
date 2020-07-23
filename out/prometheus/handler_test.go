package prometheus

import (
	"testing"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/id"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ray := xray.New(xray.NewSyncEmitter, id.Generator20Base64)
	p := NewExporter(func(args []xray.Arg) (out []xray.Arg) {
		for _, a := range args {
			if a.Name() == "id" {
				out = append(out, a)
			}
		}
		return
	},
		nil,
		100*time.Millisecond,
		1*time.Second,
	)
	ray.On(p.Handle)

	ray.Inc("foo", args.ID64(10))
	ray.Inc("bar")
	ray.Increment("foo", 5, args.ID64(10))
	ray.Inc("foo", args.ID64(11))

	if assert.Len(t, p.counters, 3) {
		assert.Equal(t, int64(1), p.counters["foo\tid\t11"].value)
		assert.Equal(t, int64(6), p.counters["foo\tid\t10"].value)
		assert.Equal(t, int64(1), p.counters["bar\t"].value)
	}

	ray.Gauge("foo", 100, args.ID64(10))
	ray.Gauge("bar", 100)
	ray.Gauge("foo", 3, args.ID64(10))

	if assert.Len(t, p.gauges, 2) {
		assert.Equal(t, int64(3), p.gauges["foo\tid\t10"].value)
		assert.Equal(t, int64(100), p.gauges["bar\t"].value)
	}

	ray.Duration("foo", 50*time.Microsecond, args.ID64(3))
	ray.Duration("foo", 15*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 50*time.Millisecond, args.ID64(42))
	ray.Duration("foo", 100*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 500*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 1000*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 2000*time.Millisecond, args.ID64(3))

	if assert.Len(t, p.histogram, 2) {
		assert.Equal(t, int64(6), p.histogram["foo\tid\t3"].counts)
		assert.Equal(t, int64(3615050000), p.histogram["foo\tid\t3"].sum)
		if assert.Len(t, p.histogram["foo\tid\t3"].buckets, 3) {
			assert.Equal(t, int64(3), p.histogram["foo\tid\t3"].buckets[0])
			assert.Equal(t, int64(2), p.histogram["foo\tid\t3"].buckets[1])
			assert.Equal(t, int64(1), p.histogram["foo\tid\t3"].buckets[2])
		}
	}
}

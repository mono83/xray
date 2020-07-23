package prometheus

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/mono83/xray/id"
	"github.com/stretchr/testify/assert"
)

var expectedWriteOutput = `# TYPE bar gauge
bar 100
# TYPE foo gauge
foo{id="10"} 3
# TYPE bar counter
bar 1
# TYPE foo counter
foo{id="11"} 1
# TYPE foo counter
foo{id="10"} 6
`

func TestWrite(t *testing.T) {
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

	ray.Inc("foo", args.ID64(10), args.Name("Wuut"))
	ray.Inc("bar")
	ray.Increment("foo", 5, args.ID64(10))
	ray.Inc("foo", args.ID64(11))
	ray.Gauge("foo", 100, args.ID64(10))
	ray.Gauge("bar", 100)
	ray.Gauge("foo", 3, args.ID64(10))
	ray.Duration("foo", 50*time.Microsecond, args.ID64(3))
	ray.Duration("foo", 15*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 50*time.Millisecond, args.ID64(42))
	ray.Duration("foo", 100*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 500*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 1000*time.Millisecond, args.ID64(3))
	ray.Duration("foo", 2000*time.Millisecond, args.ID64(3))

	bts := bytes.NewBuffer(nil)
	if assert.NoError(t, p.Write(bts)) {
		fmt.Println(bts.String())
		assert.Equal(t, expectedWriteOutput, bts.String())
	}
}

func TestEscape(t *testing.T) {
	assert.Equal(t, string(escape("foo-bar.baz")), "foo_bar_baz")
	assert.Equal(t, string(escape("hello world")), "hello_world")
	assert.Equal(t, string(escape(" trim spaces  ")), "trim_spaces")
}

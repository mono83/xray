package influxdb

import (
	"bytes"
	"github.com/mono83/xray"
	"github.com/mono83/xray/text"
	"strconv"
)

// NewBuffer builds and returns buffer, used to write metrics event is InfluxDB UDP format
// If filter function is provided, than in will be used to determine argument, that
// are allowed to send
// If nil provided, arguments will be ignored
func NewBuffer(filter xray.ArgFilter) *Buffer {
	b := &Buffer{Buffer: bytes.NewBuffer(nil)}
	if filter != nil {
		b.argsAllowed = true
		b.argFilter = filter
	}

	return b
}

// Buffer is special type of buffer, used to build StatsD-compatible packets.
type Buffer struct {
	*bytes.Buffer
	argsAllowed   bool           // True if params printing in Dogstats format allowed
	argWasWritten bool           // True if params output was started at current line
	argFilter     xray.ArgFilter // Function, used for argument filtering
}

// WriteEvent writes event data into bytes buffer
func (b *Buffer) WriteEvent(event xray.MetricsEvent) {
	b.WriteString(event.GetKey())
	if b.argsAllowed {
		for _, arg := range b.argFilter(event.Args()) {
			n, v := text.SanitizeArg(arg)
			b.WriteRune(',')
			b.Write(n)
			b.WriteRune('=')
			b.Write(v)
		}
	}

	b.WriteString(" value=")
	b.WriteString(strconv.FormatInt(event.GetValue(), 10))
	b.WriteRune('\n')
}

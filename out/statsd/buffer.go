package statsd

import (
	"bytes"
	"github.com/mono83/xray"
	"github.com/mono83/xray/text"
	"strconv"
)

// NewBuffer builds and returns buffer, used to write metrics event is StatsD format
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
	count         int            // Amount of placed events
	argsAllowed   bool           // True if params printing in Dogstats format allowed
	argWasWritten bool           // True if params output was started at current line
	argFilter     xray.ArgFilter // Function, used for argument filtering
}

// WriteEvent writes event data into bytes buffer
func (b *Buffer) WriteEvent(event xray.MetricsEvent) {
	if b.count > 0 {
		b.WriteRune('\n')
	}
	b.WriteString(event.GetKey())
	b.WriteRune(':')
	switch event.GetType() {
	case xray.GAUGE:
		b.WriteString(strconv.FormatInt(event.GetValue(), 10))
		b.WriteRune('|')
		b.WriteRune('g')
	case xray.DURATION:
		b.WriteString(strconv.FormatInt(event.GetValue(), 10))
		b.WriteString("|ms")
	default:
		b.WriteString(strconv.FormatInt(event.GetValue(), 10))
		b.WriteRune('|')
		b.WriteRune('c')
	}

	args := event.Args()
	if b.argsAllowed && len(args) > 0 {
		for _, param := range b.argFilter(args) {
			b.WriteArg(param)
		}
	}
	b.count++
}

// WriteArg writes argument (if allowed) to buffer
func (b *Buffer) WriteArg(arg xray.Arg) {
	if arg != nil {
		if b.argWasWritten {
			b.WriteRune(',')
		} else {
			b.WriteString("|@1.0|#")
			b.argWasWritten = true
		}

		n, v := text.SanitizeArg(arg)

		b.Write(n)
		b.WriteRune(':')
		b.Write(v)
	}
}

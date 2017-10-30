package statsd

import (
	"bytes"
	"github.com/mono83/xray"
	"io"
	"strconv"
	"strings"
)

type sender struct {
	target     io.Writer
	argAllowed bool
	argFilter  xray.ArgFilter
}

func (s sender) handle(events ...xray.Event) {
	builder := &packetBuilder{
		Buffer:      bytes.NewBuffer(nil),
		argsAllowed: s.argAllowed,
		argFilter:   s.argFilter,
	}

	for _, event := range events {
		if m, ok := event.(xray.MetricsEvent); ok {
			builder.WriteEvent(m)
		}

	}

	if builder.count > 0 {
		s.target.Write(builder.Bytes())
	}
}

// packetBuilder is special type of buffer, used to build StatsD-compatible packets.
type packetBuilder struct {
	*bytes.Buffer
	count         int            // Amount of placed events
	argsAllowed   bool           // True if params printing in Dogstats format allowed
	argWasWritten bool           // True if params output was started at current line
	argFilter     xray.ArgFilter // Function, used for argument filtering
}

// WriteEvent writes event data into bytes buffer
func (pb *packetBuilder) WriteEvent(event xray.MetricsEvent) {
	if pb.count > 0 {
		pb.WriteRune('\n')
	}
	pb.WriteString(event.GetKey())
	pb.WriteRune(':')
	switch event.GetType() {
	case xray.GAUGE:
		pb.WriteString(strconv.FormatInt(event.GetValue(), 10))
		pb.WriteRune('|')
		pb.WriteRune('g')
	case xray.DURATION:
		pb.WriteString(strconv.FormatInt(event.GetValue(), 10))
		pb.WriteString("|ms")
	default:
		pb.WriteString(strconv.FormatInt(event.GetValue(), 10))
		pb.WriteRune('|')
		pb.WriteRune('c')
	}

	args := event.Args()
	if pb.argsAllowed && len(args) > 0 {
		for _, param := range pb.argFilter(args) {
			pb.WriteArg(param)
		}
	}
	pb.count++
}

// WriteArg writes argument (if allowed) to buffer
func (pb *packetBuilder) WriteArg(arg xray.Arg) {
	if arg != nil {
		if pb.argWasWritten {
			pb.WriteRune(',')
		} else {
			pb.WriteString("|@1.0|#")
			pb.argWasWritten = true
		}

		pb.WriteString(arg.Name())
		pb.WriteRune(':')
		pb.Write(SanitizeParamValue(arg.Value()))
	}
}

var sanitizeReplacement = byte('_')

// SanitizeParamValue replaces special characters from param value
func SanitizeParamValue(value string) []byte {
	if len(value) == 0 {
		return []byte{}
	}

	bts := []byte(strings.TrimSpace(value))
	for i, v := range bts {
		if !(v == 46 || (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			bts[i] = sanitizeReplacement
		}
	}

	return bts
}

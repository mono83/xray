package statsd

import (
	"github.com/mono83/xray"
	"io"
)

type sender struct {
	target     io.Writer
	argAllowed bool
	argFilter  xray.ArgFilter
}

func (s sender) handle(events ...xray.Event) {
	var buf *Buffer
	if s.argAllowed {
		buf = NewBuffer(s.argFilter)
	} else {
		buf = NewBuffer(nil)
	}

	for _, event := range events {
		if m, ok := event.(xray.MetricsEvent); ok {
			buf.WriteEvent(m)
		}

	}

	if buf.count > 0 {
		s.target.Write(buf.Bytes())
	}
}

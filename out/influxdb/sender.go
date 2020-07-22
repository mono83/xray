package influxdb

import (
	"io"

	"github.com/mono83/xray"
)

type sender struct {
	target     io.Writer
	argAllowed bool
	argFilter  xray.ArgFilter
}

func (s sender) handle(events ...xray.Event) {
	for _, event := range events {
		s.one(event)
	}
}

func (s sender) one(event xray.Event) {
	if m, ok := event.(xray.MetricsEvent); ok {
		var buf *Buffer
		if s.argAllowed {
			buf = NewBuffer(s.argFilter)
		} else {
			buf = NewBuffer(nil)
		}
		buf.WriteEvent(m)
		s.target.Write(buf.Bytes())
	}
}

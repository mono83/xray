package logstash

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/text"
)

type sender struct {
	target    io.Writer
	argFilter xray.ArgFilter
}

func (s sender) handle(events ...xray.Event) {
	for _, event := range events {
		l, ok := event.(xray.LogEvent)
		if !ok {
			continue
		}

		pkt := map[string]interface{}{}

		if l.Size() > 0 {
			for _, arg := range s.argFilter(l.Args()) {
				pkt[arg.Name()] = arg.Scalar()
			}
		}

		pkt["object"] = l.GetLogger()
		pkt["log-level"] = text.LevelToString(l.GetLevel())
		pkt["pattern"] = l.GetMessage()
		pkt["message"] = text.InterpolatePlainText(l.GetMessage(), l, false)
		pkt["event-time"] = l.GetTime().Format(time.RFC3339)

		bts, _ := json.Marshal(pkt)
		s.target.Write(bts)
	}
}

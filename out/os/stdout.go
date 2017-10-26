package os

import (
	"github.com/mono83/xray"
	"github.com/mono83/xray/out"
	"github.com/mono83/xray/out/writer"
	"github.com/mono83/xray/text/color"
	"os"
)

// StdOutLogger returns new asynchronous events handler, that prints logs
// into standard output stream
func StdOutLogger(level xray.Level) xray.Handler {
	return out.Filter(
		out.Channel(
			writer.New(
				os.Stdout,
				color.FormatLogEvent,
			),
		),
		func(e xray.Event) bool {
			l, ok := e.(xray.LogEvent)
			return ok && l.GetLevel() >= level
		},
	)
}

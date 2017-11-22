package os

import (
	"os"

	"github.com/mono83/xray"
	"github.com/mono83/xray/out"
	"github.com/mono83/xray/out/writer"
	"github.com/mono83/xray/text/color"
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

// StdOutDefaultLogger returns new asynchronous events handler, that prints logs
// into standard output stream with following policy:
// - For BOOT and ROOT loggers prints INFO and higher levels
// - For other logger prints WARNING and higher levels
func StdOutDefaultLogger() xray.Handler {
	return out.Filter(
		out.Channel(
			writer.New(
				os.Stdout,
				color.FormatLogEvent,
			),
		),
		func(e xray.Event) bool {
			l, ok := e.(xray.LogEvent)
			return ok && (l.GetLevel() >= xray.WARNING || l.GetLogger() == "ROOT" || l.GetLogger() == "BOOT")
		},
	)
}

// StdOutDumper returns new asynchronous events handler, that prints dumping events
// into standard output stream
func StdOutDumper() xray.Handler {
	return out.Filter(
		out.Channel(
			writer.New(
				os.Stdout,
				color.FormatDumpEvent,
			),
		),
		func(e xray.Event) bool {
			_, ok := e.(xray.ByteDumpEvent)
			return ok
		},
	)
}

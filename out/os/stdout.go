package os

import (
	"os"

	"github.com/mono83/xray"
	"github.com/mono83/xray/out"
	"github.com/mono83/xray/out/writer"
	"github.com/mono83/xray/text/color"
	"github.com/mono83/xray/text/plain"
)

// stdLogger constructs standard output logger
func stdLogger(predicate func(event xray.Event) bool, colors, stderr bool) xray.Handler {
	output := os.Stdout
	if stderr {
		output = os.Stderr
	}
	formatter := plain.FormatLogEvent
	if colors {
		formatter = color.FormatLogEvent
	}

	return out.Filter(
		out.Channel(writer.New(output, formatter)),
		predicate,
	)
}

// StdOutLogger returns new asynchronous events handler, that prints logs
// into standard output stream
func StdOutLogger(level xray.Level) xray.Handler {
	return StdOutLoggerX(level, true, false)
}

// StdOutLoggerX returns new asynchronous events handler, that prints logs
// into standard output stream
// X stands for eXtended configuration
func StdOutLoggerX(level xray.Level, colors, stderr bool) xray.Handler {
	return stdLogger(
		func(e xray.Event) bool {
			l, ok := e.(xray.LogEvent)
			return ok && l.GetLevel() >= level
		},
		colors,
		stderr,
	)
}

// StdOutDefaultLogger returns new asynchronous events handler, that prints logs
// into standard output stream with following policy:
// - For BOOT and ROOT loggers prints INFO and higher levels
// - For other logger prints WARNING and higher levels
func StdOutDefaultLogger() xray.Handler {
	return StdOutDefaultLoggerX(true, false)
}

// StdOutDefaultLoggerX returns new asynchronous events handler, that prints logs
// into standard output stream with following policy:
// - For BOOT and ROOT loggers prints INFO and higher levels
// - For other logger prints WARNING and higher levels
// X stands for eXtended configuration
func StdOutDefaultLoggerX(colors, stderr bool) xray.Handler {
	return stdLogger(
		func(e xray.Event) bool {
			l, ok := e.(xray.LogEvent)
			return ok && (l.GetLevel() >= xray.WARNING || (l.GetLevel() >= xray.INFO && (l.GetLogger() == "ROOT" || l.GetLogger() == "BOOT")))
		},
		colors,
		stderr,
	)
}

// StdOutDumper returns new asynchronous events handler, that prints dumping events
// into standard output stream
func StdOutDumper() xray.Handler {
	return StdOutDumperX(true, false)
}

// StdOutDumperX returns new asynchronous events handler, that prints dumping events
// into standard output stream
// X stands for eXtended configuration
func StdOutDumperX(colors, stderr bool) xray.Handler {
	output := os.Stdout
	if stderr {
		output = os.Stderr
	}
	formatter := plain.FormatDumpEvent
	if colors {
		formatter = color.FormatDumpEvent
	}

	return out.Filter(
		out.Channel(writer.New(output, formatter)),
		func(e xray.Event) bool {
			_, ok := e.(xray.ByteDumpEvent)
			return ok
		},
	)
}

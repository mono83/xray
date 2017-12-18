package os

import (
	"errors"
	"os"

	"github.com/mono83/xray"
	"github.com/mono83/xray/out"
	"github.com/mono83/xray/out/writer"
	"github.com/mono83/xray/text/plain"
)

// LogFile returns logger, that will write messages into file
func LogFile(filename string, level xray.Level) (xray.Handler, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if _, ok := err.(*os.PathError); ok {
		// Trying to create
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, errors.New(filename + " is folder")
	}

	return out.Filter(
		out.Channel(
			writer.New(
				file,
				plain.FormatLogEvent,
			),
		),
		func(e xray.Event) bool {
			l, ok := e.(xray.LogEvent)
			return ok && l.GetLevel() >= level
		},
	), nil
}

// LogAndDumpFile returns logger, that will write messages and dump info into file
func LogAndDumpFile(filename string, level xray.Level) (xray.Handler, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if _, ok := err.(*os.PathError); ok {
		// Trying to create
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, errors.New(filename + " is folder")
	}

	return out.FanOut(
		out.Filter(
			out.Channel(
				writer.New(
					file,
					plain.FormatLogEvent,
				),
			),
			func(e xray.Event) bool {
				l, ok := e.(xray.LogEvent)
				return ok && l.GetLevel() >= level
			},
		),
		out.Filter(
			out.Channel(
				writer.New(
					file,
					plain.FormatDumpEvent,
				),
			),
			func(e xray.Event) bool {
				_, ok := e.(xray.ByteDumpEvent)
				return ok
			},
		),
	), nil
}

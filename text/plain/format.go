package plain

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/mono83/xray"
	"github.com/mono83/xray/text"
)

// FormatLogEvent performs plaintext formatting of log event
func FormatLogEvent(e xray.Event) string {
	if e == nil {
		return ""
	}

	l, ok := e.(xray.LogEvent)
	if !ok {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(text.TimeFormat(l.GetTime()))
	buf.WriteRune(' ')
	if l.GetLevel() == xray.ALERT {
		buf.WriteString(" ALERT ")
		buf.WriteRune(' ')
	}
	if l.GetLevel() == xray.CRITICAL {
		buf.WriteString(" CRIT ")
		buf.WriteRune(' ')
	}
	buf.WriteString(text.Interpolate(l.GetMessage(), e, func(arg xray.Arg) string {
		return text.PlainInterpolator(arg)
	}))
	if logger := l.GetLogger(); len(logger) > 0 {
		buf.WriteRune(' ')
		buf.WriteString("@" + strings.ToLower(logger))
	}

	return buf.String()
}

// FormatDumpEvent performs plaintext formatting of log event
func FormatDumpEvent(e xray.Event) string {
	if e == nil {
		return ""
	}

	d, ok := e.(xray.ByteDumpEvent)
	if !ok {
		return ""
	}

	bts := d.GetBytes()

	buf := bytes.NewBuffer(nil)
	buf.WriteString(text.TimeFormat(d.GetTime()))
	switch d.GetSource() {
	case xray.OUT:
		buf.WriteString(" >>> ")
	case xray.IN:
		buf.WriteString(" <<< ")
	}
	buf.WriteString("Dump contents (")
	buf.WriteString(strconv.Itoa(len(bts)))
	buf.WriteString(" bytes)")
	if len(bts) > 0 {
		buf.WriteRune('\n')
		buf.WriteString(string(bts))
	}

	return buf.String()
}

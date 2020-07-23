package color

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/mono83/xray"
	"github.com/mono83/xray/text"
)

// FormatLogEvent performs colorful formatting of log event
func FormatLogEvent(e xray.Event) string {
	if e == nil {
		return ""
	}

	l, ok := e.(xray.LogEvent)
	if !ok {
		return ""
	}

	loggerColor := getRayColor(l.GetLogger())
	rayColor := getRayColor(l.GetRayID())
	mainColor := getMessageColor(l.GetLevel())
	varColorResolver := getVarColorResolver(l.GetLevel())

	buf := bytes.NewBuffer(nil)
	buf.WriteString(loggerColor.Sprint("▐"))
	buf.WriteString(rayColor.Sprint("▌"))
	buf.WriteString(colorTime.Sprint(text.TimeFormat(l.GetTime())))
	buf.WriteRune(' ')
	if l.GetLevel() == xray.ALERT {
		buf.WriteString(colorBadgeAlert.Sprint(" ALRT "))
		buf.WriteRune(' ')
	}
	if l.GetLevel() == xray.CRITICAL {
		buf.WriteString(colorBadgeCrit.Sprint(" CRIT "))
		buf.WriteRune(' ')
	}
	buf.WriteString(mainColor.Sprint(text.Interpolate(l.GetMessage(), e, func(arg xray.Arg) string {
		return varColorResolver(arg).Sprint(text.PlainInterpolator(arg)) + mainColor.Open()
	})))
	if logger := l.GetLogger(); len(logger) > 0 {
		buf.WriteRune(' ')
		buf.WriteString(colorLogger.Sprint("@" + strings.ToLower(logger)))
	}

	return buf.String()
}

// FormatDumpEvent performs colorful formatting of log event
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
	buf.WriteString(colorTime.Sprint(text.TimeFormat(d.GetTime())))
	switch d.GetSource() {
	case xray.OUT:
		buf.WriteString(colorDumpType.Sprint(" >>> "))
	case xray.IN:
		buf.WriteString(colorDumpType.Sprint(" <<< "))
	}
	buf.WriteString(colorDebug.Sprint("Dump contents ("))
	buf.WriteString(colorDebug.Sprint(strconv.Itoa(len(bts))))
	buf.WriteString(colorDebug.Sprint(" bytes)"))
	if len(bts) > 0 {
		buf.WriteRune('\n')
		buf.WriteString(colorDebug.Sprint(string(bts)))
	}

	return buf.String()
}

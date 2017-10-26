package color

import (
	"bytes"
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

	mainColor := getMessageColor(l.GetLevel())
	varColorResolver := getVarColorResolver(l.GetLevel())

	buf := bytes.NewBuffer(nil)
	buf.WriteString(colorTime.Sprint(text.TimeFormat(l.GetTime())))
	buf.WriteRune(' ')
	if l.GetLevel() == xray.ALERT {
		buf.WriteString(colorBadge.Sprint(" ALERT "))
		buf.WriteRune(' ')
	}
	if l.GetLevel() == xray.CRITICAL {
		buf.WriteString(colorBadge.Sprint(" CRIT "))
		buf.WriteRune(' ')
	}
	buf.WriteString(mainColor.Sprint(text.Interpolate(l.GetMessage(), e, func(arg xray.Arg) string {
		return varColorResolver(arg).Sprint(text.PlainInterpolator(arg)) + mainColor.Open()
	})))

	return buf.String()
}

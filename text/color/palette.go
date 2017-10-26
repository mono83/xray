package color

import (
	"github.com/mono83/xray"
)

var colorTime = New(FgHiBlack)
var colorDebug = New(FgHiBlack)
var colorInfo = New(FgGreen)
var colorWarning = New(FgYellow)
var colorError = New(FgRed)
var colorVarLow = New(FgHiBlack, Underline)
var colorVarNormal = New(FgCyan)
var colorVarHi = New(FgYellow)
var colorBadge = New(BgRed, FgHiWhite, Bold)
var colorDumpType = New(FgCyan)

func getMessageColor(level xray.Level) Color {
	switch level {
	case xray.TRACE, xray.DEBUG:
		return colorDebug
	case xray.INFO:
		return colorInfo
	case xray.WARNING:
		return colorWarning
	default:
		return colorError
	}
}

func getVarColorResolver(level xray.Level) func(a xray.Arg) Color {
	return func(a xray.Arg) Color {
		if level < xray.INFO {
			return colorVarLow
		} else if level > xray.WARNING {
			return colorVarHi
		}

		return colorVarNormal
	}
}

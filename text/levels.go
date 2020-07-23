package text

import (
	"strings"

	"github.com/mono83/xray"
)

// LevelToString returns string representation of xray Level
func LevelToString(l xray.Level) string {
	switch l {
	case xray.TRACE:
		return "trace"
	case xray.DEBUG:
		return "debug"
	case xray.INFO:
		return "info"
	case xray.WARNING:
		return "warning"
	case xray.ERROR:
		return "error"
	case xray.ALERT:
		return "alert"
	case xray.CRITICAL:
		return "critical"
	default:
		return "unknown"
	}
}

// ParseLevel parses string into xray Level
func ParseLevel(v string) xray.Level {
	switch strings.ToLower(v) {
	case "trace":
		return xray.TRACE
	case "debug":
		return xray.DEBUG
	case "info":
		return xray.INFO
	case "warn", "warning":
		return xray.WARNING
	case "error":
		return xray.ERROR
	case "alert":
		return xray.ALERT
	case "critical", "emergency":
		return xray.CRITICAL
	default:
		return xray.INFO
	}
}

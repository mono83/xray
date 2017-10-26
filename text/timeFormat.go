package text

import "time"

// TimeFormat provides default time formatting
func TimeFormat(t time.Time) string {
	return t.Format("02 15:04:05.000000")
}

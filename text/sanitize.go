package text

import (
	"strings"

	"github.com/mono83/xray"
)

var sanitizeReplacement = byte('_')

// Sanitize function takes string value and removes values, that can be unsafe
// for remote receivers. In fact, it leaves only alpha numeric values, others are
// replaced with underscore
func Sanitize(value string) []byte {
	if len(value) == 0 {
		return []byte{}
	}

	bts := []byte(strings.TrimSpace(value))
	for i, v := range bts {
		if !(v == 46 || (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			bts[i] = sanitizeReplacement
		}
	}

	return bts
}

// SanitizeArg returns sanitized version of argument key and value
func SanitizeArg(a xray.Arg) (key []byte, value []byte) {
	key = Sanitize(a.Name())
	value = Sanitize(a.Value())
	return
}

package xray

import (
	"github.com/mono83/xray/args"
	"regexp"
)

// placeholdersRegex contains rules to find placeholders inside string
var placeholdersRegex = regexp.MustCompile(":[0-9a-zA-Z\\-_]+")

// Interpolate replaces all placeholders within source string using arguments bucket
// and string formatter
func Interpolate(source string, bucket Bucket, format func(Arg) string) string {
	if len(source) <= 1 || bucket == nil || bucket.Size() == 0 || format == nil {
		return source
	}

	return placeholdersRegex.ReplaceAllStringFunc(
		source,
		func(x string) string {
			arg := bucket.Get(x[1:])
			if arg == nil {
				arg = args.Nil(x[1:])
			}

			return format(arg)
		},
	)
}

func plainInterpolatorCommon(a Arg) string {
	if a == nil {
		return ""
	}
	return a.Value()
}

func plainInterpolatorBracketed(a Arg) string {
	if a == nil {
		return "<!>"
	} else if _, ok := a.(args.Nil); ok {
		return "<!" + a.Name() + "!>"
	}

	return "[" + a.Value() + "]"
}

// InterpolatePlainText performs plaintext interpolation
func InterpolatePlainText(source string, bucket Bucket, brackets bool) string {
	if brackets {
		return Interpolate(source, bucket, plainInterpolatorBracketed)
	}

	return Interpolate(source, bucket, plainInterpolatorCommon)
}

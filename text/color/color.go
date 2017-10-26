package color

// This is small port of glorious color lib https://github.com/fatih/color

import (
	"fmt"
	"strconv"
	"strings"
)

// Color represents ANSI color sequence
type Color []Attribute

// New builds new color using
func New(attrs ...Attribute) Color {
	return Color(attrs)
}

// sequence returns a formated SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c Color) sequence() string {
	params := []Attribute(c)
	format := make([]string, len(params))
	for i, v := range params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// Sprint performs color formatting
func (c Color) Sprint(a ...interface{}) string {
	return c.Open() + fmt.Sprint(a...) + c.Close()
}

// Open returns open ANSI sequence
func (c Color) Open() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

// Close returns resetting ANSI sequence
func (Color) Close() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

package writer

import (
	"fmt"
	"github.com/mono83/xray"
	"io"
)

// New builds new events handler, used to print to arbitrary writer
func New(w io.Writer, eventFormat func(xray.Event) string) xray.Handler {
	return func(events ...xray.Event) {
		for _, e := range events {
			if str := eventFormat(e); len(str) > 0 {
				fmt.Fprintln(w, str)
			}
		}
	}
}

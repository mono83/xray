package std

import (
	"github.com/mono83/xray"
	"github.com/mono83/xray/out/os"
)

// Bootstrap performs logger bootstrapping
// Not for production usage
func Bootstrap(level xray.Level, printDumps bool) {
	xray.ROOT.On(os.StdOutLogger(level))
	if printDumps {
		xray.ROOT.On(os.StdOutDumper())
	}
}

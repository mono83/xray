package xcobra

import (
	"os"
	"time"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/spf13/cobra"
)

// Start starts application command
func Start(cmd *cobra.Command) {
	if err := Wrap(cmd).Execute(); err != nil {
		xray.ROOT.Alert("Application done with error :err", args.Error{Err: err})
		os.Exit(1)
	}

	// Log delivery timeout
	time.Sleep(100 * time.Millisecond)
}

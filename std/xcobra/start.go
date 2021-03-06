package xcobra

import (
	"os"

	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"github.com/spf13/cobra"
)

// Start starts application command
func Start(cmd *cobra.Command, postWrapperCallbacks ...func(*cobra.Command)) {
	c := Wrap(cmd)

	if len(postWrapperCallbacks) > 0 {
		for _, clb := range postWrapperCallbacks {
			clb(c)
		}
	}

	err := c.Execute()
	code := 0

	if err != nil {
		xray.ROOT.Alert("Application done with error :err", args.Error{Err: err})
		code = 1
	}

	// Wait log delivery
	xray.Waiter.Wait()

	os.Exit(code)
}

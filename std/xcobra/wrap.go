package xcobra

import (
	"github.com/mono83/xray"
	"github.com/mono83/xray/out/os"
	"github.com/spf13/cobra"
)

// Wrap method wraps spf13 cobra command, injecting logging control options
func Wrap(cmd *cobra.Command) *cobra.Command {
	if cmd == nil {
		return cmd
	}

	if cmd.PersistentFlags().Lookup("verbose") == nil {
		cmd.PersistentFlags().BoolP("verbose", "v", false, "Display info-level logging and higher")
	}
	if cmd.PersistentFlags().Lookup("vv") == nil {
		cmd.PersistentFlags().Bool("vv", false, "Very verbose mode, debug will be displayed")
	}
	if cmd.PersistentFlags().Lookup("vvv") == nil {
		cmd.PersistentFlags().Bool("vvv", false, "Extra verbose mode, trace and debug will be displayed")
	}
	if cmd.PersistentFlags().Lookup("quiet") == nil {
		cmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet mode, logging output will be suppressed")
	}
	if cmd.PersistentFlags().Lookup("no-ansi") == nil {
		cmd.PersistentFlags().Bool("no-ansi", false, "Disable ANSI coloring for logs")
	}

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		vv, _ := cmd.Flags().GetBool("vv")
		vvv, _ := cmd.Flags().GetBool("vvv")
		verbose, _ := cmd.Flags().GetBool("verbose")
		quiet, _ := cmd.Flags().GetBool("quiet")
		//nocolor, _ := cmd.Flags().GetBool("no-ansi") TODO

		// Enabling logger
		if !quiet {
			if vvv {
				// Extra verbose mode
				xray.ROOT.On(os.StdOutLogger(xray.TRACE))
			} else if vv {
				// Very verbose mode
				xray.ROOT.On(os.StdOutLogger(xray.DEBUG))
			} else if verbose {
				// Info+ logging
				xray.ROOT.On(os.StdOutLogger(xray.INFO))
			} else {
				// Default logging - warning & higher + logs from BOOT and ROOT
				xray.ROOT.On(os.StdOutDefaultLogger())
			}

		}
	}

	return cmd
}

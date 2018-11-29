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
	if cmd.PersistentFlags().Lookup("vd") == nil {
		cmd.PersistentFlags().Bool("vd", false, "Verbose dump mode, will dump packet contents, if app supports it")
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
	if cmd.PersistentFlags().Lookup("stderr") == nil {
		cmd.PersistentFlags().Bool("stderr", false, "Use STDERR instead of STDOUT for logging")
	}

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		vv, _ := cmd.Flags().GetBool("vv")
		vd, _ := cmd.Flags().GetBool("vd")
		vvv, _ := cmd.Flags().GetBool("vvv")
		verbose, _ := cmd.Flags().GetBool("verbose")
		quiet, _ := cmd.Flags().GetBool("quiet")
		noColor, _ := cmd.Flags().GetBool("no-ansi")
		useStderr, _ := cmd.Flags().GetBool("stderr")

		// Enabling logger
		if !quiet {
			if vvv {
				// Extra verbose mode
				xray.ROOT.On(os.StdOutLoggerX(xray.TRACE, !noColor, useStderr))
			} else if vv {
				// Very verbose mode
				xray.ROOT.On(os.StdOutLoggerX(xray.DEBUG, !noColor, useStderr))
			} else if verbose {
				// Info+ logging
				xray.ROOT.On(os.StdOutLoggerX(xray.INFO, !noColor, useStderr))
			} else {
				// Default logging - warning & higher + logs from BOOT and ROOT
				xray.ROOT.On(os.StdOutDefaultLoggerX(!noColor, useStderr))
			}
			if vd {
				// Packet dump mode
				xray.ROOT.On(os.StdOutDumperX(!noColor, useStderr))
			}
		}
	}

	return cmd
}

// Package prepare contains the implementation of the prepare subcommand
package prepare

import (
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/medley"
)

var PrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare everything needed to test shutter.",
	Run:   medley.ShowHelpAndExit,
}

func init() {
	PrepareCmd.AddCommand(configCmd)
	PrepareCmd.AddCommand(fundCmd)

	initConfigFlags()
	initFundFlags()
}

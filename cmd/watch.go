package cmd

import (
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches and verifies every new block's commitment proposed by a mev-boost proposer",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

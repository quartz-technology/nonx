package cmd

import (
	"github.com/quartz-technology/charon/watch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Verifies if the commitments hold for every new proposed blocks",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Extract the required configuration for the watch command.
		configuration := watch.ConfigurationFromViper(viper.GetViper())

		// Runs the watch loop which will verify new mev-boost proposed payloads.
		return watch.Run(configuration)
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
}

package cmd

import (
	"github.com/quartz-technology/charon/verify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies if a commitment holds for a single slot.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Builds the configuration which holds the clients and the desired slot to analyze.
		configuration := verify.ConfigurationFromViper(viper.GetViper())

		// Run the verification for the given slot.
		if err := verify.Run(configuration); err != nil {
			return err
		}

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	verify.Flags(viper.GetViper(), verifyCmd.Flags())
}

package cmd

import (
	"github.com/0xpanoramix/frd-go/constants"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "charon",
	Short: "A verifier for mev-boost commitments sent by proposers to relays.",
	Long: `
🪬 Charon uses the Flashbots Data Transparency API to verify if a proposer broke its commitments
to use a relay's given payload.

🤖 It is able to listen to new mev-boost blocks as they are proposed to the network or inspect a
single block.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Registering specific root flags.
	rootFlags(viper.GetViper(), rootCmd.Flags())
	// Registering the configuration flags, mapped to the environment and to a configuration file.
	configurationFlags(viper.GetViper(), rootCmd.Flags())

	// Add the subcommand verify.
	rootCmd.AddCommand(verifyCmd)
	// Add the subcommand watch.
	rootCmd.AddCommand(watchCmd)
}

var configurationFilePath string

func initConfig() {
	if configurationFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configurationFilePath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".charon-config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".charon-config")
	}

	viper.SetEnvPrefix("CHARON")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

const (
	endpointExecutionClientNodeFlag     = "endpoint-node"
	endpointExecutionClientNodeViperKey = "endpoint.node"
	endpointExecutionClientNodeEnv      = "CHARON_ENDPOINT_NODE"

	endpointRelayFlag     = "endpoint-relay"
	endpointRelayViperKey = "endpoint.relay"
	endpointRelayEnv      = "CHARON_ENDPOINT_RELAY"
)

func rootFlags(_ *viper.Viper, f *pflag.FlagSet) {
	// --config
	rootCmd.PersistentFlags().StringVar(&configurationFilePath, "config", "", "path to the configuration file")
}

func configurationFlags(v *viper.Viper, f *pflag.FlagSet) {
	// --endpoint-node
	f.String(endpointExecutionClientNodeFlag, "http://localhost:8545", "The execution client JSON-RPC URL.")
	err := v.BindPFlag(endpointExecutionClientNodeViperKey, f.Lookup(endpointExecutionClientNodeFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(endpointExecutionClientNodeViperKey, endpointExecutionClientNodeEnv)
	cobra.CheckErr(err)

	// --endpoint-relay
	f.String(endpointRelayFlag, constants.FlashbotsRelayMainnet, "The relay endpoint to query the Data Transparency API.")
	err = v.BindPFlag(endpointRelayViperKey, f.Lookup(endpointRelayFlag))
	cobra.CheckErr(err)
	err = v.BindEnv(endpointRelayViperKey, endpointRelayEnv)
	cobra.CheckErr(err)
}

type charonConfig struct {
	executionClientNodeEndpoint string
	relayEndpoint               string
}

func charonConfigFromViper(v *viper.Viper) *charonConfig {
	configuration := &charonConfig{}

	configuration.executionClientNodeEndpoint = v.GetString(endpointExecutionClientNodeViperKey)
	configuration.relayEndpoint = v.GetString(endpointRelayViperKey)

	return configuration
}

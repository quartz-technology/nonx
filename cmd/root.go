package cmd

import (
	"github.com/quartz-technology/charon/root"
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
	rootCmd.PersistentFlags().StringVar(&configurationFilePath, "config", "", "path to the configuration file")

	// Registering the configuration flags, mapped to the environment and to a configuration file.
	root.Flags(viper.GetViper(), rootCmd.Flags())

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

package root

import (
	"github.com/0xpanoramix/frd-go/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	endpointExecutionClientNodeFlag     = "endpoint-node"
	endpointExecutionClientNodeViperKey = "endpoint.node"
	endpointExecutionClientNodeEnv      = "NONX_ENDPOINT_NODE"

	endpointRelayFlag     = "endpoint-relay"
	endpointRelayViperKey = "endpoint.relay"
	endpointRelayEnv      = "NONX_ENDPOINT_RELAY"
)

func Flags(v *viper.Viper, f *pflag.FlagSet) {
	configurationFlags(v, f)
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

func GetEndpointExecutionClient(v *viper.Viper) string {
	return v.GetString(endpointExecutionClientNodeViperKey)
}

func GetEndpointRelay(v *viper.Viper) string {
	return v.GetString(endpointRelayViperKey)
}

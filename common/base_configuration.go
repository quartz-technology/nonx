package common

import (
	"github.com/0xpanoramix/frd-go/data"
	"github.com/quartz-technology/charon/root"
	"github.com/spf13/viper"
	"time"
)

// BaseConfiguration holds the clients able to query the Relay Data Transparency API and the
// Beacon Chain.
type BaseConfiguration struct {
	DC *data.TransparencyClient
	EC *EthClient
}

// BaseConfigurationFromViper creates a new BaseConfiguration using the bound cobra flags and
// viper keys.
func BaseConfigurationFromViper(v *viper.Viper) *BaseConfiguration {
	return &BaseConfiguration{
		DC: data.NewTransparencyClient(root.GetEndpointRelay(v), time.Second),
		EC: NewEthClient(root.GetEndpointExecutionClient(v)),
	}
}

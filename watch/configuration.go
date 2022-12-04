package watch

import (
	"github.com/0xpanoramix/frd-go/data"
	"github.com/quartz-technology/charon/common"
	"github.com/quartz-technology/charon/root"
	"github.com/spf13/viper"
	"time"
)

type Configuration struct {
	dc *data.TransparencyClient
	ec *common.EthClient
}

func ConfigurationFromViper(v *viper.Viper) *Configuration {
	return &Configuration{
		dc: data.NewTransparencyClient(root.GetEndpointRelay(v), time.Second),
		ec: common.NewEthClient(root.GetEndpointExecutionClient(v)),
	}
}

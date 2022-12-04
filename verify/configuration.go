package verify

import (
	"github.com/quartz-technology/charon/common"
	"github.com/spf13/viper"
)

// Configuration holds the data transparency and beacon clients as well as the desired slot to
// analyze.
type Configuration struct {
	base *common.BaseConfiguration
	slot uint64
}

// ConfigurationFromViper creates a new Configuration using the base configuration extractor and
// the value for the bound --slot flag and its corresponding viper key.
func ConfigurationFromViper(v *viper.Viper) *Configuration {
	return &Configuration{
		base: common.BaseConfigurationFromViper(v),
		slot: GetSlotToVerify(v),
	}
}

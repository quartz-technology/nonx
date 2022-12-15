package watch

import (
	"github.com/quartz-technology/nonx/common"
	"github.com/spf13/viper"
)

// Configuration is used by the watch Run method and holds the common configuration only,
// as it does not require any extra configuration to perform the analysis for now.
type Configuration struct {
	base *common.BaseConfiguration
}

// ConfigurationFromViper creates a new Configuration using the base configuration extractor.
func ConfigurationFromViper(v *viper.Viper) *Configuration {
	return &Configuration{
		base: common.BaseConfigurationFromViper(v),
	}
}

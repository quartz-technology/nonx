package watch

import (
	"github.com/quartz-technology/charon/common"
	"github.com/spf13/viper"
)

type Configuration struct {
	base *common.BaseConfiguration
}

func ConfigurationFromViper(v *viper.Viper) *Configuration {
	return &Configuration{
		base: common.BaseConfigurationFromViper(v),
	}
}

package pkg

import (
	"github.com/spf13/viper"
)

type GengConfig struct {
}

// GetConfig get's configuration from viper
func GetConfig[T any]() *T {
	logger := GetLogger()

	var cfg T
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatal("cannot unmarshal configuration object", "err", err)
	}

	return &cfg
}

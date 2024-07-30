package pkg

import "github.com/spf13/viper"

type GengConfig struct {
}

func NewConfig() (*GengConfig, error) {
	var cfg GengConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

  return &cfg, nil
}

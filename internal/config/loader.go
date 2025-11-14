package config

import "github.com/spf13/viper"

const cfgPath = "./config/config.yaml"

func Load() (*Config, error) {
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

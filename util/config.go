// Package util holds general utility function for the application
package util

import (
	"github.com/spf13/viper"
)

type InfrastructureConfig struct {
	ConfigName string `mapstructure:"CONFIG_NAME"`
}

func LoadInfrastructureConfig() (config InfrastructureConfig, err error) {

	// Config the env file values
	viper.SetConfigName("app-dev")
	viper.SetConfigType("env")

	// Overwrite values from file if env values exist
	viper.AutomaticEnv()

	// Read the config from file or env
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	viper.WatchConfig()

	return
}

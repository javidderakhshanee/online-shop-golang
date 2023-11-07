package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigurationsWrapper struct {
	MySql MySqlConfigurations
}

type MySqlConfigurations struct {
	ConnectionString string
}

type Configuration struct {
	ConfigurationsWrapper
}

func NewConfiguration() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration ConfigurationsWrapper

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return Configuration{configuration}
}

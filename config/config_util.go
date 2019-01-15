package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	config Config
)

func init()  {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func Get() Config {
	return config
}
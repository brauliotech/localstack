package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config Settings

type Settings struct {
	Address string
	Region  string
	Profile string
	ID      string
	Secret  string
}

func init() {
	InitConfig()
}

func InitConfig() {
	viper.SetDefault("Address", "http://localhost:4566")
	viper.SetDefault("region", "us-east-1")
	viper.SetDefault("Profiele", "dev")
	viper.SetDefault("ID", "test")
	viper.SetDefault("Secret", "test")

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}
}

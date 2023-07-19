package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get environment variable %v", err)
	}
	Config.AddConfigPath(wd)
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("Could not get yaml : %v", err)
	}
}

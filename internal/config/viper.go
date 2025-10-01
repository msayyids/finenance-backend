package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")

	if err := config.ReadInConfig(); err != nil {
		log.Printf("No .env file found, skip... (%v)", err)
	}

	config.AutomaticEnv()
	return config

}

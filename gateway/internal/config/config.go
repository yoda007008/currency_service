package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		host string
		port string
	}
	Auth struct {
		URL string
	}
	CurrencyService struct {
		URL string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName(".yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}

}

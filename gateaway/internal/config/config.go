package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	CurrencyServiceUrl string
	GatewayPort        string
}

func LoadConfig(filepath string) (*Config, error) {
	if err := godotenv.Load(filepath); err != nil {
		return nil, err
	}

	return &Config{
		CurrencyServiceUrl: os.Getenv("CURRENCY_SERVICE_URL"), // добавить в .env путь до currency
		GatewayPort:        os.Getenv("GATEWAY_PORT"),
	}, nil
}

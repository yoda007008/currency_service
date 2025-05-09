package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Auth struct {
		URL string `yaml:"url"`
	} `yaml:"auth"`

	CurrencyService struct {
		URL string `yaml:"url"`
	} `yaml:"currency-service"`

	SecretKey struct {
		Token string `yaml:"token"`
	} `yaml:"secret_key"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type (
	ApiConfig struct {
		ApiPort string
	}

	Config struct {
		ApiConfig
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.ApiPort == "" {
		return errors.New("enviroment required")
	}

	return nil
}

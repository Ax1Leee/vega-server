package config

import (
	"github.com/spf13/viper"

	"errors"
)

type Config struct {
	// Embedding
	*viper.Viper
}

func LoadConfig(path string) (*Config, error) {
	// Initialize config
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	// Read config
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.New("failed to load config")
	}
	return &Config{Viper: v}, nil
}

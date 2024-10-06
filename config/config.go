package config

import (
	"fmt"
	"github-actions-nginx-demo/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Logger *logger.Config `mapstructure:"logger"`
	GRPC   *GRPC          `mapstructure:"grpc"`
}

type GRPC struct {
	Port string `mapstructure:"port" validate:"required"`
}

func LoadConfig(pathToConfig string) (*Config, error) {
	if pathToConfig == "" {
		return nil, fmt.Errorf("path to cfg is empty")
	}

	cfg := &Config{}

	viper.SetConfigType("yml")
	viper.SetConfigFile(pathToConfig)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("(ReadInConfig)")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cfg")
	}

	return cfg, nil
}

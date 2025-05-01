package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Env      string `mapstructure:"ENV" validate:"required"`
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.Load(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg *Config) Load() error {
	cfg.Env = os.Getenv("ENV")

	switch cfg.Env {
	case "dev":
		if err := loadDevConfig(cfg); err != nil {
			return err
		}
	case "test":
		if err := loadTestConfig(cfg); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid ENV value: %s", cfg.Env)
	}

	if err := validateConfig(cfg); err != nil {
		return err
	}

	return nil
}

func loadDevConfig(cfg *Config) error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err)
	}
	return nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return fmt.Errorf("configuration validation failed: %v", err)
	}
	return nil
}

func loadTestConfig(cfg *Config) error {
	viper.SetConfigFile("test.env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err)
	}
	return nil
}

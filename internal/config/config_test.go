package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	// Setting up environment variable for testing
	t.Setenv("ENV", "dev")
	t.Setenv("HTTP_PORT", "8080")

	t.Run("valid config - test environment", func(t *testing.T) {
		t.Setenv("ENV", "test")
		cfg, err := NewConfig()
		require.NoError(t, err)
		assert.Equal(t, "test", cfg.Env)
	})

	t.Run("invalid ENV value", func(t *testing.T) {
		t.Setenv("ENV", "invalid")
		cfg, err := NewConfig()
		require.Error(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("missing HTTP_PORT", func(t *testing.T) {
		// Remove HTTP_PORT environment variable
		t.Setenv("ENV", "dev")
		t.Setenv("HTTP_PORT", "")
		cfg, err := NewConfig()
		require.Error(t, err)
		assert.Nil(t, cfg)
	})
}

func TestLoadDevConfig(t *testing.T) {
	t.Run("failed to load dev config", func(t *testing.T) {
		// Simulate an error while loading .env file
		t.Setenv("ENV", "dev")
		viper.SetConfigFile(".invalid_env")
		cfg := &Config{}
		err := cfg.Load()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error reading config file")
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		cfg := &Config{
			Env:      "dev",
			HTTPPort: "8080",
		}
		err := validateConfig(cfg)
		require.NoError(t, err)
	})

	t.Run("invalid config", func(t *testing.T) {
		cfg := &Config{
			Env:      "", // Missing Env
			HTTPPort: "8080",
		}
		err := validateConfig(cfg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "configuration validation failed")
	})
}

func TestLoadTestConfig(t *testing.T) {
	t.Run("valid test config", func(t *testing.T) {
		t.Setenv("ENV", "test")
		t.Setenv("HTTP_PORT", "8081")

		cfg := &Config{}
		err := cfg.Load()
		require.NoError(t, err)

		assert.Equal(t, "test", cfg.Env)
		assert.Equal(t, "8081", cfg.HTTPPort)
	})
}

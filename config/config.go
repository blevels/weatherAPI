// Package config utilizes the viper package to initialize application configuration information from the config.yml file.
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App     `yaml:"app"`
		HTTP    `yaml:"http"`
		Weather `yaml:"weather"`
		Log     `yaml:"log"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Weather struct {
		Key string `env-required:"true" yaml:"key"   env:"API_KEY"`
		URI string `env-required:"true" yaml:"uri"   env:"API_URI"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error unable to read the config file: %v", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error unable to decode config: %v", err)
	}

	return cfg, nil
}

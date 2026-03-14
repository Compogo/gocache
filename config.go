package gocache

import (
	"time"

	"github.com/Compogo/compogo/configurator"
)

const (
	ExpirationFieldName      = "cache.memory.expiration"
	CleanupIntervalFieldName = "cache.memory.cleanup"

	ExpirationDefault      = 5 * time.Minute
	CleanupIntervalDefault = 10 * time.Minute
)

type Config struct {
	Expiration      time.Duration
	CleanupInterval time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.Expiration == 0 || config.Expiration == ExpirationDefault {
		configurator.SetDefault(ExpirationFieldName, ExpirationDefault)
		config.Expiration = configurator.GetDuration(ExpirationFieldName)
	}

	if config.CleanupInterval == 0 || config.CleanupInterval == CleanupIntervalDefault {
		configurator.SetDefault(CleanupIntervalFieldName, CleanupIntervalDefault)
		config.CleanupInterval = configurator.GetDuration(CleanupIntervalFieldName)
	}

	return config
}

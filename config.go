package gocache

import (
	"time"

	"github.com/Compogo/compogo"
)

const (
	// ExpirationFieldName время жизни данных
	ExpirationFieldName = "cache.memory.expiration"

	// CleanupIntervalFieldName интервал очистки просроченных данных
	CleanupIntervalFieldName = "cache.memory.cleanup"
)

var (
	// ExpirationDefault время жизни по умолчанию
	ExpirationDefault = 5 * time.Minute

	// CleanupIntervalDefault интервал очистки по умолчанию
	CleanupIntervalDefault = 10 * time.Minute
)

// Config содержит конфигурацию in-memory кэша.
type Config struct {
	Expiration      time.Duration
	CleanupInterval time.Duration
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
// Если значения не заданы, устанавливаются значения по умолчанию.
func Configuration(config *Config, configurator compogo.Configurator) *Config {
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

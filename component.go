package gocache

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	goCache "github.com/patrickmn/go-cache"
)

// Component is a ready-to-use Compogo component that provides an in-memory cache.
// It automatically:
//   - Registers Config and Cache in the DI container
//   - Adds command-line flags for expiration and cleanup interval
//   - Configures the cache during Configuration phase
//   - Provides the cache as both *goCache.Cache and Cache interface
//
// Usage:
//
//	compogo.WithComponents(
//	    gocache.Component,
//	    // ... your service components
//	)
//
// Then in your service:
//
//	type Service struct {
//	    cache gocache.Cache  // or *goCache.Cache
//	}
var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewCache,
			func(cache *goCache.Cache) Cache { return cache },
		)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.DurationVar(
				&config.Expiration,
				ExpirationFieldName,
				ExpirationDefault,
				"default data retention time",
			)

			flagSet.DurationVar(
				&config.CleanupInterval,
				CleanupIntervalFieldName,
				CleanupIntervalDefault,
				"period for cleaning the cache from 'rotten' data",
			)
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}

package registration

import (
	"github.com/Compogo/cache"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/gocache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/eko/gocache/store/go_cache/v4"
)

// Component is a ready-to-use Compogo component that registers the in-memory cache
// driver with the central cache system.
//
// It depends on gocache.Component to ensure the underlying in-memory cache is
// initialized. The actual registration happens in init(), which is safe because
// cache.Registration only stores factory functions without requiring runtime state.
var Component = &component.Component{
	Dependencies: component.Components{
		gocache.Component,
	},
}

// init registers the "memory" cache driver with the central cache system.
// The registration happens at program startup, independent of component lifecycle.
//
// The factory function receives a container and:
//   - Extracts the cache configuration (*cache.Config) and the gocache client
//   - Creates a gocache.Store (compatible with store.StoreInterface)
//   - Returns it to the cache system for wrapping with metrics
//
// This driver can be selected with --cache.driver=memory.
func init() {
	cache.Registration("memory", func(container container.Container) (store.StoreInterface, error) {
		var cacheStore store.StoreInterface
		var err error

		err = container.Invoke(func(config *cache.Config, client gocache.Cache) {
			cacheStore = go_cache.NewGoCache(client, store.WithExpiration(config.Expiration))
		})

		if err != nil {
			return nil, err
		}

		return cacheStore, nil
	})
}

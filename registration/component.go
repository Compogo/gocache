package registration

import (
	"github.com/Compogo/cache"
	"github.com/Compogo/compogo"
	"github.com/Compogo/gocache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/eko/gocache/store/go_cache/v4"
)

// Component — компонент регистрации in-memory драйвера для gocache.
// Регистрирует драйвер "memory" в системе кэширования.
//
// После подключения этого компонента, пакет cache сможет использовать
// in-memory хранилище как бекенд для кэширования.
//
// Пример:
//
//	app.AddComponents(
//	    &registration.Component, // регистрация драйвера для cache
//	)
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&gocache.Component,
	},
}

// Регистрация драйвера "memory" в системе cache.
// Использует gocache.Cache как источник данных.
func init() {
	cache.Registration("memory", func(container compogo.Container) (store.StoreInterface, error) {
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

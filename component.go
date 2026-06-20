package gocache

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	goCache "github.com/patrickmn/go-cache"
)

// Component — компонент in-memory кэша для Compogo.
// Регистрирует конфигурацию и экземпляр кэша в DI-контейнере.
//
// Кэш хранится в памяти приложения и подходит для:
//   - Кэширования небольших объёмов данных
//   - Временного хранения сессий
//   - Кэширования результатов вычислений
//
// Пример:
//
//	app.AddComponents(&gocache.Component)
//
//	var c gocache.Cache
//	container.Invoke(func(cache gocache.Cache) { c = cache })
//	c.Set("key", "value", time.Minute)
var Component = compogo.Component{
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			NewConfig,
			NewCache,
			func(cache *goCache.Cache) Cache { return cache },
		)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
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
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}

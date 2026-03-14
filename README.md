# Compogo GoCache 🗃️

**Compogo GoCache** — это in-memory кэш для Compogo, построенный поверх популярной библиотеки [patrickmn/go-cache](https://github.com/patrickmn/go-cache). Предоставляет полный API оригинальной библиотеки, настраивается через флаги и может использоваться как самостоятельный кэш или как драйвер для централизованной системы кэширования [Compogo Cache](https://github.com/Compogo/cache).

## 🚀 Установка

```bash
go get github.com/Compogo/gocache
```

### 📦 Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/gocache"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        gocache.Component,  // добавляем in-memory кэш
        compogo.WithComponents(
            userServiceComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}

// Использование в сервисе
var userServiceComponent = &component.Component{
    Dependencies: component.Components{gocache.Component},
    Execute: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(cache gocache.Cache) {
            service := &UserService{cache: cache}
            service.Start()
        })
    }),
}

type UserService struct {
    cache gocache.Cache
}

func (s *UserService) GetUser(id int) (*User, error) {
    // Пытаемся достать из кэша
    if data, found := s.cache.Get(fmt.Sprintf("user:%d", id)); found {
        return data.(*User), nil
    }
    
    // Нет в кэше — грузим из БД
    user, err := s.db.LoadUser(id)
    if err != nil {
        return nil, err
    }
    
    // Кладём в кэш
    s.cache.Set(fmt.Sprintf("user:%d", id), user, cache.DefaultExpiration)
    
    return user, nil
}
```

### ✨ Возможности

#### 🎯 Полный API go-cache

Интерфейс `Cache` повторяет все методы оригинальной библиотеки:

```go
// Базовые операции
cache.Set("key", value, 5*time.Minute)
value, found := cache.Get("key")
cache.Delete("key")

// Атомарные инкременты для всех типов
cache.IncrementInt("counter", 1)
cache.DecrementFloat64("metric", 0.5)

// Работа с expiration
value, expTime, found := cache.GetWithExpiration("key")
cache.SetDefault("key", value) // использует TTL по умолчанию

// Очистка
cache.DeleteExpired()
cache.Flush()

// Персистентность
cache.SaveFile("cache.backup")
cache.LoadFile("cache.backup")

// Обработка событий
cache.OnEvicted(func(key string, value interface{}) {
    log.Printf("evicted: %s", key)
})
```

#### 🔌 Два способа использования

##### Как самостоятельный кэш:

```go
type Service struct {
    cache gocache.Cache  // через интерфейс
    // или
    raw *goCache.Cache    // напрямую
}
```

##### Как драйвер для централизованной системы кэширования:

```go
import (
    "github.com/Compogo/cache"
    "github.com/Compogo/gocache/registration"
)

app := compogo.NewApp("myapp",
    cache.Component,
    gocache.Component,
    registration.Component,  // регистрирует "memory" драйвер
)

// Теперь можно использовать через cache.CacheInterface[[]byte]
// с флагом --cache.driver=memory
```

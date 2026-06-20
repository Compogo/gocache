# Compogo GoCache

In-memory кэш для фреймворка [Compogo](https://github.com/Compogo/compogo).

На основе [patrickmn/go-cache](https://github.com/patrickmn/go-cache) предоставляет:

* Хранение данных в памяти приложения
* Автоматическую очистку просроченных данных
* Инкременты и декременты числовых значений
* Сохранение/загрузку кэша в файл
* Колбэк при удалении данных

## Установка

```shell
go get github.com/Compogo/gocache
```

## Конфигурация

### Флаги командной строки

```shell
# Время жизни данных по умолчанию
--cache.memory.expiration=5m

# Интервал очистки просроченных данных
--cache.memory.cleanup=10m
```

## Интеграция с [Compogo Cache](https://github.com/Compogo/cache)

Этот пакет также регистрирует драйвер `memory` в системе кэширования [Compogo Cache](https://github.com/Compogo/cache):

```go
import (
    "github.com/Compogo/gocache/registration"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithComponents(&registration.Component),
    )
    // Теперь cache.NewCache() использует in-memory хранилище
}
```

## Зависимости

* [Compogo Cache](https://github.com/Compogo/cache) — основной фреймворк
* [patrickmn/go-cache](https://github.com/patrickmn/go-cache) — библиотека in-memory кэша

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

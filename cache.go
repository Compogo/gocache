package gocache

import (
	"io"
	"time"

	"github.com/Compogo/compogo"
	goCache "github.com/patrickmn/go-cache"
)

// Cache — интерфейс для in-memory кэша на основе go-cache.
// Предоставляет методы для работы с кэшем в памяти приложения.
//
// Поддерживает:
//   - Сохранение данных с TTL
//   - Инкременты и декременты числовых значений
//   - Сохранение/загрузку в файл
//   - Колбэк на удаление данных
//   - Очистку просроченных данных
//
// Пример:
//
//	var c gocache.Cache
//	container.Invoke(func(cache gocache.Cache) { c = cache })
//
//	c.Set("user:123", User{Name: "John"}, time.Minute)
//	val, found := c.Get("user:123")
type Cache interface {
	Set(string, interface{}, time.Duration)
	SetDefault(string, interface{})
	Add(string, interface{}, time.Duration) error
	Replace(string, interface{}, time.Duration) error
	Get(string) (interface{}, bool)
	GetWithExpiration(string) (interface{}, time.Time, bool)
	Increment(string, int64) error
	IncrementFloat(string, float64) error
	IncrementInt(string, int) (int, error)
	IncrementInt8(string, int8) (int8, error)
	IncrementInt16(string, int16) (int16, error)
	IncrementInt32(string, int32) (int32, error)
	IncrementInt64(string, int64) (int64, error)
	IncrementUint(string, uint) (uint, error)
	IncrementUintptr(string, uintptr) (uintptr, error)
	IncrementUint8(string, uint8) (uint8, error)
	IncrementUint16(string, uint16) (uint16, error)
	IncrementUint32(string, uint32) (uint32, error)
	IncrementUint64(string, uint64) (uint64, error)
	IncrementFloat32(string, float32) (float32, error)
	IncrementFloat64(string, float64) (float64, error)
	Decrement(string, int64) error
	DecrementFloat(string, float64) error
	DecrementInt(string, int) (int, error)
	DecrementInt8(string, int8) (int8, error)
	DecrementInt16(string, int16) (int16, error)
	DecrementInt32(string, int32) (int32, error)
	DecrementInt64(string, int64) (int64, error)
	DecrementUint(string, uint) (uint, error)
	DecrementUintptr(string, uintptr) (uintptr, error)
	DecrementUint8(string, uint8) (uint8, error)
	DecrementUint16(string, uint16) (uint16, error)
	DecrementUint32(string, uint32) (uint32, error)
	DecrementUint64(string, uint64) (uint64, error)
	DecrementFloat32(string, float32) (float32, error)
	DecrementFloat64(string, float64) (float64, error)
	Delete(string)
	DeleteExpired()
	OnEvicted(f func(string, interface{}))
	Save(io.Writer) error
	SaveFile(string) error
	Load(r io.Reader) error
	LoadFile(string) error
	Items() map[string]goCache.Item
	ItemCount() int
	Flush()
}

// NewCache создаёт новый in-memory кэш.
// Принимает конфигурацию (TTL и интервал очистки) и логгер.
func NewCache(config *Config, logger compogo.Logger) *goCache.Cache {
	logger = logger.GetLogger("cache").GetLogger("inAppMemory")

	logger.Infof("expiration - %s", config.Expiration)
	logger.Infof("cleanupInterval - %s", config.CleanupInterval)

	return goCache.New(config.Expiration, config.CleanupInterval)
}

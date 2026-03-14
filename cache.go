package gocache

import (
	"io"
	"time"

	"github.com/Compogo/compogo/logger"
	goCache "github.com/patrickmn/go-cache"
)

// Cache defines the complete interface for in-memory cache operations.
// It mirrors the full API of patrickmn/go-cache to provide maximum flexibility.
//
// The interface includes methods for:
//   - Basic operations: Set, Get, Delete
//   - Atomic increments/decrements for all numeric types
//   - Expiration management
//   - Persistence (Save/Load to/from files)
//   - Event handling (OnEvicted)
//   - Cache inspection (Items, ItemCount)
type Cache interface {
	Set(k string, x interface{}, d time.Duration)
	SetDefault(k string, x interface{})
	Add(k string, x interface{}, d time.Duration) error
	Replace(k string, x interface{}, d time.Duration) error
	Get(k string) (interface{}, bool)
	GetWithExpiration(k string) (interface{}, time.Time, bool)
	Increment(k string, n int64) error
	IncrementFloat(k string, n float64) error
	IncrementInt(k string, n int) (int, error)
	IncrementInt8(k string, n int8) (int8, error)
	IncrementInt16(k string, n int16) (int16, error)
	IncrementInt32(k string, n int32) (int32, error)
	IncrementInt64(k string, n int64) (int64, error)
	IncrementUint(k string, n uint) (uint, error)
	IncrementUintptr(k string, n uintptr) (uintptr, error)
	IncrementUint8(k string, n uint8) (uint8, error)
	IncrementUint16(k string, n uint16) (uint16, error)
	IncrementUint32(k string, n uint32) (uint32, error)
	IncrementUint64(k string, n uint64) (uint64, error)
	IncrementFloat32(k string, n float32) (float32, error)
	IncrementFloat64(k string, n float64) (float64, error)
	Decrement(k string, n int64) error
	DecrementFloat(k string, n float64) error
	DecrementInt(k string, n int) (int, error)
	DecrementInt8(k string, n int8) (int8, error)
	DecrementInt16(k string, n int16) (int16, error)
	DecrementInt32(k string, n int32) (int32, error)
	DecrementInt64(k string, n int64) (int64, error)
	DecrementUint(k string, n uint) (uint, error)
	DecrementUintptr(k string, n uintptr) (uintptr, error)
	DecrementUint8(k string, n uint8) (uint8, error)
	DecrementUint16(k string, n uint16) (uint16, error)
	DecrementUint32(k string, n uint32) (uint32, error)
	DecrementUint64(k string, n uint64) (uint64, error)
	DecrementFloat32(k string, n float32) (float32, error)
	DecrementFloat64(k string, n float64) (float64, error)
	Delete(k string)
	DeleteExpired()
	OnEvicted(f func(string, interface{}))
	Save(w io.Writer) (err error)
	SaveFile(fname string) error
	Load(r io.Reader) error
	LoadFile(fname string) error
	Items() map[string]goCache.Item
	ItemCount() int
	Flush()
}

// NewCache creates a new in-memory cache instance.
// It initializes the underlying go-cache with the configured expiration and cleanup interval.
// The cache starts its cleanup goroutine automatically.
//
// The informer is used to log the configuration for debugging purposes.
func NewCache(config *Config, informer logger.Informer) *goCache.Cache {
	informer.Infof("[cache.app] expiration - %s", config.Expiration)
	informer.Infof("[cache.app] cleanupInterval - %s", config.CleanupInterval)

	return goCache.New(config.Expiration, config.CleanupInterval)
}

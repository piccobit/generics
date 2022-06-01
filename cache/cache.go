/*
Package cache is a simple generic implementation of a cache.
*/
package cache

import (
	"sync"
)

type Cache[T any] struct {
	content map[string]T
	maxSize int
	mutex   sync.RWMutex
}

type UnderflowError struct{}
type OverflowError struct{}

func (ue *UnderflowError) Error() string {
	return "Underflow error"
}

func (ue *OverflowError) Error() string {
	return "Overflow error"
}

// New returns the pointer to a new cache.
// The 'maxSize' parameter allows to specify a
// maximum size for the cache. Setting this to 0
// allows the cache to grow infinitely.
func New[T any](maxSize int) *Cache[T] {
	cache := Cache[T]{
		content: map[string]T{},
		maxSize: maxSize,
	}

	return &cache
}

// Load tries to get a cached value from the provided key.
// The returned boolean value indicates if the operation was
// successful or not.
func (p *Cache[T]) Load(key string) (T, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if value, ok := p.content[key]; ok {
		return value, true
	}

	var ret T

	return ret, false
}

// Save stores the given value indexed by the also provided key,
// If the maximum size of the cache is reached an Overflow error
// is returned.
func (p *Cache[T]) Save(key string, value T) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.maxSize > 0 && len(p.content) >= p.maxSize {
		return &OverflowError{}
	}

	p.content[key] = value

	return nil
}

/*
Package lfucache is a simple generic implementation of an LFU (Least Frequently Used) cache.
To solve the fact that an entry which has only been used in the beginning is not dropped later,
the frequency counters are divided by 2 before a new entry is added.
*/
package lfucache

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type item[T any] struct {
	freq  int
	added time.Time
	value T
}

type LFUCache[T any] struct {
	content       map[string]item[T]
	maxSize       int
	idDropItem    string
	hitsCounter   uint
	missedCounter uint
	mutex         sync.RWMutex
}

type UnderflowError struct{}
type OverflowError struct{}
type DuplicateError struct{}

type IDInterface interface {
	ID() string
}

func (e *UnderflowError) Error() string {
	return "Underflow error"
}

func (e *OverflowError) Error() string {
	return "Overflow error"
}

func (e *DuplicateError) Error() string {
	return "Duplicate error"
}

// New returns the pointer to a new LFU cache.
// The 'maxSize' parameter allows to specify a
// maximum size for the stack.
func New[T any](maxSize int) *LFUCache[T] {
	cache := LFUCache[T]{
		content: make(map[string]item[T]),
		maxSize: maxSize,
	}

	return &cache
}

// Get returns the value stored by the provided ID.
// If the ID doesn't exist 'false' is returned.
func (p *LFUCache[T]) Get(id string) (T, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var ok bool

	if ok = p.contains(id); !ok {
		var dummy T

		p.missedCounter++

		return dummy, false
	}

	p.hitsCounter++

	cacheItem := p.content[id]

	cacheItem.freq++

	p.content[id] = cacheItem

	return p.content[id].value, true
}

// Contains checks if the cache contains an element with
// the provided ID.
func (p *LFUCache[T]) Contains(id string) bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if _, ok := p.content[id]; ok {
		return true
	}

	return false
}

// contains checks if the cache contains an element with
// the provided ID.
// This function is only used internally and does not use
// the mutex to lock during the read access.
func (p *LFUCache[T]) contains(id string) bool {
	if _, ok := p.content[id]; ok {
		return true
	}

	return false
}

// String implements the Stringer interface.
func (p *LFUCache[T]) String() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var str strings.Builder

	str.WriteString("[")

	followingItems := false

	for _, cacheItem := range p.content {
		if followingItems {
			str.WriteString(",")
		} else {
			followingItems = true
		}

		_, _ = fmt.Fprintf(&str, "%v", cacheItem.value)
	}

	str.WriteString("]")

	return str.String()
}

// AddByID adds the provided argument with the provided ID to the LFU cache.
// If the added item is a new one and the LFU cache has already reached its
// maximum size, the oldest item is dropped to make place for the new one.
// If the added item is already part of the LFU cache it will be moved to the
// end of the cache.
func (p *LFUCache[T]) AddByID(id string, arg T) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.reduceFrequency()

	if ok := p.contains(id); !ok {
		if len(p.content) >= p.maxSize {
			p.dropLFU()
		}

		cacheItem := item[T]{
			value: arg,
			added: time.Now(),
		}

		p.content[id] = cacheItem
		p.idDropItem = id
	} else {
		return &DuplicateError{}
	}

	return nil
}

// Add adds the provided argument to the LFU cache.
// The ID used is either provided using the ID interface or generated internally.
// If the added item is a new one and the LFU cache has already reached its
// maximum size, the oldest item is dropped to make place for the new one.
// If the added item is already part of the LFU cache it will be moved to the
// end of the cache.
func (p *LFUCache[T]) Add(arg T) (string, error) {
	var id string

	if idInterface, ok := any(arg).(IDInterface); ok {
		id = idInterface.ID()
	} else {
		id = uuid.New().String()
	}

	return id, p.AddByID(id, arg)
}

func (p *LFUCache[T]) dropLFU() {
	minFreq := 0

	for id, cacheItem := range p.content {
		if cacheItem.freq > minFreq {
			continue
		} else if cacheItem.freq == minFreq {
			timeDropItem := p.content[p.idDropItem].added
			timeCurrentItem := p.content[id].added

			if timeCurrentItem.After(timeDropItem) {
				continue
			} else if timeCurrentItem.Equal(timeDropItem) {
				continue
			}

			p.idDropItem = id
			minFreq = cacheItem.freq
		}
	}

	delete(p.content, p.idDropItem)
}

// Stats returns the statistics of the provided cache,
// the hits & the misses.
func (p *LFUCache[T]) Stats() (uint, uint) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.hitsCounter, p.missedCounter
}

func (p *LFUCache[T]) reduceFrequency() {
	for _, cacheItem := range p.content {
		cacheItem.freq = cacheItem.freq / 2
	}
}

// GetCache returns the cache content so that it can be
// used in a 'for range' loop.
func (p *LFUCache[T]) GetCache() map[string]T {
	content := make(map[string]T)
	for k, v := range p.content {
		content[k] = v.value
	}

	return content
}

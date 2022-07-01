/*
Package lrucache is a simple generic implementation of an LRU (Last Recently Used) cache.
*/
package lrucache

import (
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type item[T any] struct {
	id    string
	value T
}

type LRUCache[T any] struct {
	content []item[T]
	maxSize int
	mutex   sync.RWMutex
}

type UnderflowError struct{}
type OverflowError struct{}

type IDInterface interface {
	ID() string
}

func (e *UnderflowError) Error() string {
	return "Underflow error"
}

func (e *OverflowError) Error() string {
	return "Overflow error"
}

// New returns the pointer to a new LRU cache.
// The 'maxSize' parameter allows to specify a
// maximum size for the stack.
func New[T any](maxSize int) *LRUCache[T] {
	cache := LRUCache[T]{maxSize: maxSize}

	return &cache
}

// Get returns the value stored by the provided ID.
// If the ID doesn't exist 'false' is returned.
func (p *LRUCache[T]) Get(id string) (T, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var idx int
	var ok bool

	if idx, ok = p.contains(id); !ok {
		var dummy T

		return dummy, false
	}

	return p.content[idx].value, true
}

// Contains checks if the cache contains an element with
// the provided ID.
func (p *LRUCache[T]) Contains(id string) (int, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for idx, i := range p.content {
		if i.id == id {
			return idx, true
		}
	}

	return -1, false
}

// contains checks if the cache contains an element with
// the provided ID.
// This function is only used internally and does not use
// the mutex to lock during the read access.
func (p *LRUCache[T]) contains(id string) (int, bool) {
	for idx, i := range p.content {
		if i.id == id {
			return idx, true
		}
	}

	return -1, false
}

// String implements the Stringer interface.
func (p *LRUCache[T]) String() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var str strings.Builder

	str.WriteString("[")

	for i, cacheItem := range p.content {
		if i > 0 {
			str.WriteString(",")

		}

		_, _ = fmt.Fprintf(&str, "%v", cacheItem.value)
	}

	str.WriteString("]")

	return str.String()
}

// AddByID adds the provided argument with the provided ID to the LRU cache.
// If the added item is a new one and the LRU cache has already reached its
// maximum size, the oldest item is dropped to make place for the new one.
// If the added item is already part of the LRU cache it will be moved to the
// end of the cache.
func (p *LRUCache[T]) AddByID(id string, arg T) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if idx, ok := p.contains(id); ok {
		valueAtIndex := p.content[idx]
		before := p.content[:idx]
		after := p.content[idx+1:]
		newContent := append(before, after...)
		p.content = append(newContent, valueAtIndex)
	} else {
		if len(p.content) < p.maxSize {
			p.content = append(p.content, item[T]{id, arg})
		} else {
			newContent := p.content[1:]
			newContent = append(newContent, item[T]{id, arg})
			p.content = newContent
		}
	}

	return nil
}

// Add adds the provided argument to the LRU cache.
// The ID used is either provided using the ID interface or generated internally.
// If the added item is a new one and the LRU cache has already reached its
// maximum size, the oldest item is dropped to make place for the new one.
// If the added item is already part of the LRU cache it will be moved to the
// end of the cache.
func (p *LRUCache[T]) Add(arg T) (string, error) {
	var id string

	if idInterface, ok := any(arg).(IDInterface); ok {
		id = idInterface.ID()
	} else {
		id = uuid.New().String()
	}

	return id, p.AddByID(id, arg)
}

// GetCache returns the cache content so that it can be
// used in a 'for range' loop.
func (p *LRUCache[T]) GetCache() map[string]T {
	content := make(map[string]T)
	for _, v := range p.content {
		content[v.id] = v.value
	}

	return content
}

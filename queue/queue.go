/*
Package queue is a simple generic implementation of a FIFO ('First In, First Out') stack, that means
the first input value is the one which will be also retrieved first.
*/
package queue

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/exp/constraints"
)

type Queue[T constraints.Ordered] struct {
	content []T
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

// New returns the pointer to a new queue.
// The 'maxSize' parameter allows to specify a
// maximum size for the queue. Setting this to 0
// allows the queue to grow infinitely.
func New[T constraints.Ordered](maxSize int) *Queue[T] {
	queue := Queue[T]{maxSize: maxSize}

	return &queue
}

// Push pushes the given arguments on the provided queue.
// An overflow error is returned in case the queue is
// limited in its size and the push would overflow the queue.
func (p *Queue[T]) Push(args ...T) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.maxSize > 0 && len(p.content) >= p.maxSize {
		return &OverflowError{}
	}

	if p.maxSize > 0 && (len(p.content)+len(args)) > p.maxSize {
		return &OverflowError{}
	}

	p.content = append(p.content, args...)

	return nil
}

// String implements the Stringer interface to provide a
// textual representation of the queue content.
func (p *Queue[T]) String() string {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var str strings.Builder

	str.WriteString("[")

	for i, value := range p.content {
		if i > 0 {
			str.WriteString(",")

		}

		_, _ = fmt.Fprintf(&str, "%v", value)
	}

	str.WriteString("]")

	return str.String()
}

// Pop pops the first element of the queue and returns it to the caller.
// If the queue is empty an 'Underflow error' is returned.
func (p *Queue[T]) Pop() (T, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := p.content[0]

	p.content = p.content[1:]

	return value, nil
}

// Drop drops the first element of the queue.
// An underflow error is returned in case the queue is
// already empty.
func (p *Queue[T]) Drop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.content) <= 0 {
		return &UnderflowError{}
	}

	p.content = p.content[1:]

	return nil
}

// Length returns the number of queue elements.
func (p *Queue[T]) Length() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return len(p.content)
}

// Peek gets the first element of the queue and returns it to the caller.
// If the queue is empty an 'Underflow error' is returned.
func (p *Queue[T]) Peek() (T, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if len(p.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := p.content[0]

	return value, nil
}

// GetQueue returns the queue content so that it can be
// used in a 'for range' loop.
func (p *Queue[T]) GetQueue() []T {
	return p.content
}

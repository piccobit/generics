/*
Package queue is a simple generic implementation of a FIFO ('First In, First Out') stack, that means
the first input value is the one which will be also retrieved first.
*/
package queue

import (
	"fmt"
	"strings"
	"sync"
)

type Queue[T any] struct {
	content []T
	maxSize int
	mutex   sync.Mutex
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
func New[T any](maxSize int) *Queue[T] {
	queue := Queue[T]{maxSize: maxSize}

	return &queue
}

// Push pushes the given arguments on the provided queue.
// An overflow error is returned in case the queue is
// limited in its size and the push would overflow the queue.
func (p *Queue[T]) Push(args ...T) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	queue := *p

	if queue.maxSize > 0 && len(queue.content) >= queue.maxSize {
		return &OverflowError{}
	}

	if queue.maxSize > 0 && (len(queue.content)+len(args)) > queue.maxSize {
		return &OverflowError{}
	}

	queue.content = append(queue.content, args...)

	*p = queue

	return nil
}

// String implements the Stringer interface to provide a
// textual representation of the queue content.
func (p *Queue[T]) String() string {
	var str strings.Builder

	str.WriteString("[")

	queue := *p

	for i, value := range queue.content {
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

	queue := *p

	if len(queue.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := queue.content[0]

	queue.content = queue.content[1:]
	*p = queue

	return value, nil
}

// Drop drops the first element of the queue.
// An underflow error is returned in case the queue is
// already empty.
func (p *Queue[T]) Drop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	queue := *p

	if len(queue.content) <= 0 {
		return &UnderflowError{}
	}

	queue.content = queue.content[1:]
	*p = queue

	return nil
}

// Length returns the number of queue elements.
func (p *Queue[T]) Length() int {
	queue := *p
	return len(queue.content)
}

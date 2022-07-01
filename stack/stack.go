/*
Package stack is a simple generic implementation of a LIFO ('Last In, First Out') stack, that means
the last input value is the one which will be retrieved first.
*/
package stack

import (
	"fmt"
	"strings"
	"sync"
)

type Stack[T any] struct {
	content []T
	maxSize int
	mutex   sync.RWMutex
}

type UnderflowError struct{}
type OverflowError struct{}

func (e *UnderflowError) Error() string {
	return "Underflow error"
}

func (e *OverflowError) Error() string {
	return "Overflow error"
}

// New returns the pointer to a new stack.
// The 'maxSize' parameter allows to specify a
// maximum size for the stack. Setting this to 0
// allows the stack to grow infinitely.
func New[T any](maxSize int) *Stack[T] {
	stack := Stack[T]{maxSize: maxSize}

	return &stack
}

// Push pushes the given arguments on the provided stack.
// An overflow error is returned in case the stack is
// limited in its size and the push would overflow the stack.
func (p *Stack[T]) Push(args ...T) error {
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
// textual representation of the stack content.
func (p *Stack[T]) String() string {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

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

// Pop pops the last element of the stack and returns it to the caller.
// If the stack is empty an 'Underflow error' is returned.
func (p *Stack[T]) Pop() (T, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := p.content[len(p.content)-1]

	p.content = p.content[:len(p.content)-1]

	return value, nil
}

// Drop drops the last element of the stack.
// An underflow error is returned in case the stack is
// already empty.
func (p *Stack[T]) Drop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.content) <= 0 {
		return &UnderflowError{}
	}

	p.content = p.content[:len(p.content)-1]

	return nil
}

// Length returns the number of stack elements.
func (p *Stack[T]) Length() int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return len(p.content)
}

// Peek gets the last element of the stack and returns it to the caller.
// If the stack is empty an 'Underflow error' is returned.
func (p *Stack[T]) Peek() (T, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := p.content[len(p.content)-1]

	return value, nil
}

// GetStack returns the stack content so that it can be
// used in a 'for range' loop.
func (p *Stack[T]) GetStack() []T {
	return p.content
}

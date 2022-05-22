// Package stack is a simple implementation of a LIFO ('Last In, First Out') stack, that means
// the last input value is the one which will be retrieved first.
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

func (ue *UnderflowError) Error() string {
	return "Underflow error"
}

func (ue *OverflowError) Error() string {
	return "Overflow error"
}

// New provides the pointer to a new stack.
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
	stack := *p

	if stack.maxSize > 0 && len(stack.content) >= stack.maxSize {
		return &OverflowError{}
	}

	if stack.maxSize > 0 && (len(stack.content)+len(args)) > stack.maxSize {
		return &OverflowError{}
	}

	stack.content = append(stack.content, args...)

	*p = stack

	return nil
}

// String implements the Stringer interface to provide a
// textual representation of the stack content.
func (p *Stack[T]) String() string {
	var str strings.Builder

	str.WriteString("[")

	stack := *p

	for i, value := range stack.content {
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
	stack := *p

	if len(stack.content) <= 0 {
		var ret T
		return ret, &UnderflowError{}
	}

	value := stack.content[len(stack.content)-1]

	stack.content = stack.content[:len(stack.content)-1]
	*p = stack

	return value, nil
}

// Drop drops the last element of the stack.
// An underflow error is returned in case the stack is
// already empty.
func (p *Stack[T]) Drop() error {
	stack := *p

	if len(stack.content) <= 0 {
		return &UnderflowError{}
	}

	stack.content = stack.content[:len(stack.content)-1]
	*p = stack

	return nil
}

// Length returns the number of stack elements.
func (p *Stack[T]) Length() int {
	stack := *p
	return len(stack.content)
}

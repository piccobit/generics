package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] []T

func New[T any]() *Stack[T] {
	stack := make(Stack[T], 0)

	return &stack
}

func (p *Stack[T]) Push(args ...T) {
	stack := *p
	stack = append(stack, args...)
	*p = stack
}

func (p *Stack[T]) String() string {
	var str strings.Builder

	str.WriteString("[")

	for i, value := range *p {
		if i > 0 {
			str.WriteString(",")

		}

		fmt.Fprintf(&str, "%v", value)
	}

	str.WriteString("]")

	return str.String()
}

func (p *Stack[T]) Pop() T {
	stack := *p
	value := stack[len(stack)-1]

	stack = stack[:len(stack)-1]
	*p = stack

	return value
}

func (p *Stack[T]) Drop() {
	stack := *p
	stack = stack[:len(stack)-1]
	*p = stack

	return
}

func (p *Stack[T]) Length() int {
	stack := *p
	return len(stack)
}

package main

import (
	"fmt"
	"os"

	"generics/stack"
)

type car struct {
	name       string
	colour     string
	horsepower int
}

func main() {
	var err error

	myStringStack := stack.New[string](0)

	err = myStringStack.Push("Hello")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	err = myStringStack.Push("World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myStringStack.Length())
	fmt.Printf("Content: %v\n", myStringStack)

	stringValue, err := myStringStack.Pop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Pop: %s\n", stringValue)
	fmt.Printf("Content: %v\n", myStringStack)

	err = myStringStack.Push(stringValue, stringValue)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myStringStack)

	err = myStringStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myStringStack)

	myIntStack := stack.New[int](0)

	err = myIntStack.Push(13)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	err = myIntStack.Push(42)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myIntStack.Length())
	fmt.Printf("Content: %v\n", myIntStack)

	intValue, err := myIntStack.Pop()
	fmt.Printf("Pop: %d\n", intValue)
	fmt.Printf("Content: %v\n", myIntStack)

	err = myIntStack.Push(intValue, intValue)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myIntStack)

	err = myIntStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myIntStack)

	myCarStack := stack.New[car](0)

	err = myCarStack.Push(car{
		name:       "HD",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	err = myCarStack.Push(car{
		name:       "Doris",
		colour:     "red",
		horsepower: 59,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myCarStack.Length())
	fmt.Printf("Content: %v\n", myCarStack)

	carValue, err := myCarStack.Pop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Pop: %v\n", carValue)
	fmt.Printf("Content: %v\n", myCarStack)

	err = myCarStack.Push(carValue, carValue)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myCarStack)

	err = myCarStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myCarStack)

	myEmptyStack := stack.New[string](0)
	err = myEmptyStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	myTestOverflowStack := stack.New[string](3)

	err = myTestOverflowStack.Push("foo", "bar", "hello", "world")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}
}

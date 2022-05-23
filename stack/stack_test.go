package stack_test

import (
	"fmt"
	"os"

	"github.com/piccobit/generics/stack"
)

func ExampleStack_Push_string() {
	var err error

	myStringStack := stack.New[string](0)

	err = myStringStack.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myStringStack.Length())
	fmt.Printf("Content: %v\n", myStringStack)
	// Output:
	// Length: 2
	// Content: [Hello,World]
}

func ExampleStack_Pop_string() {
	var err error

	myStringStack := stack.New[string](0)

	err = myStringStack.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	value, err := myStringStack.Pop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Pop: %s\n", value)
	fmt.Printf("Content: %v\n", myStringStack)
	// Output:
	// Pop: World
	// Content: [Hello]
}

func ExampleStack_Drop_string() {
	var err error

	myStringStack := stack.New[string](0)

	err = myStringStack.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	err = myStringStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myStringStack)
	// Output:
	// Content: [Hello]
}

func ExampleStack_Push_int() {
	var err error

	myIntStack := stack.New[int](0)

	err = myIntStack.Push(13, 42)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myIntStack.Length())
	fmt.Printf("Content: %v\n", myIntStack)
	// Output:
	// Length: 2
	// Content: [13,42]
}

func ExampleStack_Pop_int() {
	var err error

	myIntStack := stack.New[int](0)

	err = myIntStack.Push(13, 42)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	value, err := myIntStack.Pop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Pop: %d\n", value)
	fmt.Printf("Content: %v\n", myIntStack)
	// Output:
	// Pop: 42
	// Content: [13]
}

func ExampleStack_Drop_int() {
	var err error

	myIntStack := stack.New[int](0)

	err = myIntStack.Push(13, 42)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	err = myIntStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myIntStack)
	// Output:
	// Content: [13]
}

func ExampleStack_car() {
	type car struct {
		name       string
		colour     string
		horsepower int
	}

	var err error

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
	// Output:
	// Length: 2
	// Content: [{HD blue 60},{Doris red 59}]
	// Pop: {Doris red 59}
	// Content: [{HD blue 60}]
	// Content: [{HD blue 60},{Doris red 59},{Doris red 59}]
	// Content: [{HD blue 60},{Doris red 59}]
}

func ExampleStack_underflow() {
	var err error

	myEmptyStack := stack.New[string](0)
	err = myEmptyStack.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "ERROR: %s\n", err.Error())
	}
	// Output:
	// ERROR: Underflow error
}

func ExampleStack_overflow() {
	var err error

	myTestOverflowStack := stack.New[string](3)

	err = myTestOverflowStack.Push("foo", "bar", "hello", "world")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "ERROR: %s\n", err.Error())
	}
	// Output:
	// ERROR: Overflow error
}

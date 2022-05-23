package queue_test

import (
	"fmt"
	"os"

	"github.com/piccobit/generics/queue"
)

func ExampleQueue_Push() {
	var err error

	myStringQueue := queue.New[string](0)

	err = myStringQueue.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Length: %d\n", myStringQueue.Length())
	fmt.Printf("Content: %v\n", myStringQueue)
	// Output:
	// Length: 2
	// Content: [Hello,World]
}

func ExampleQueue_Pop() {
	var err error

	myStringQueue := queue.New[string](0)

	err = myStringQueue.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	queueValue, err := myStringQueue.Pop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Pop: %s\n", queueValue)
	fmt.Printf("Content: %v\n", myStringQueue)
	// Output:
	// Pop: Hello
	// Content: [World]
}

func ExampleQueue_Drop() {
	var err error

	myStringQueue := queue.New[string](0)

	err = myStringQueue.Push("Hello", "World")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myStringQueue.Drop()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}

	fmt.Printf("Content: %v\n", myStringQueue)
	// Output:
	// Content: [World]
}

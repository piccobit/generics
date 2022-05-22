package main

import (
	"fmt"

	"generics/stack"
)

type car struct {
	name       string
	colour     string
	horsepower int
}

func main() {
	myStringStack := stack.New[string]()

	myStringStack.Push("Hello")
	myStringStack.Push("World")

	fmt.Printf("Length: %d\n", myStringStack.Length())
	fmt.Printf("Content: %v\n", myStringStack)

	stringValue := myStringStack.Pop()
	fmt.Printf("Pop: %s\n", stringValue)
	fmt.Printf("Content: %v\n", myStringStack)

	myStringStack.Push(stringValue, stringValue)
	fmt.Printf("Content: %v\n", myStringStack)

	myStringStack.Drop()
	fmt.Printf("Content: %v\n", myStringStack)

	myIntStack := stack.New[int]()

	myIntStack.Push(13)
	myIntStack.Push(42)

	fmt.Printf("Length: %d\n", myIntStack.Length())
	fmt.Printf("Content: %v\n", myIntStack)

	intValue := myIntStack.Pop()
	fmt.Printf("Pop: %d\n", intValue)
	fmt.Printf("Content: %v\n", myIntStack)

	myIntStack.Push(intValue, intValue)
	fmt.Printf("Content: %v\n", myIntStack)

	myIntStack.Drop()
	fmt.Printf("Content: %v\n", myIntStack)

	myCarStack := stack.New[car]()

	myCarStack.Push(car{
		name:       "HD",
		colour:     "blue",
		horsepower: 60,
	})
	myCarStack.Push(car{
		name:       "Doris",
		colour:     "red",
		horsepower: 59,
	})

	fmt.Printf("Length: %d\n", myCarStack.Length())
	fmt.Printf("Content: %v\n", myCarStack)

	carValue := myCarStack.Pop()
	fmt.Printf("Pop: %v\n", carValue)
	fmt.Printf("Content: %v\n", myCarStack)

	myCarStack.Push(carValue, carValue)
	fmt.Printf("Content: %v\n", myCarStack)

	myCarStack.Drop()
	fmt.Printf("Content: %v\n", myCarStack)
}

package cache_test

import (
	"fmt"

	"github.com/piccobit/generics/cache"
)

func ExampleCache_Load() {
	myCache := cache.New[string](0)

	_ = myCache.Save("foo", "foo")
	_ = myCache.Save("bar", "bar")

	var value string
	var ok bool

	value, ok = myCache.Load("foo")
	fmt.Printf("%s: %v\n", value, ok)

	value, ok = myCache.Load("bar")
	fmt.Printf("%s: %v\n", value, ok)

	value, ok = myCache.Load("foobar")
	fmt.Printf("%s: %v\n", value, ok)

	// Output:
	// foo: true
	// bar: true
	// : false
}

func ExampleCache_Save() {
	myCache := cache.New[int](0)

	_ = myCache.Save("foo", 13)
	_ = myCache.Save("bar", 42)

	var value int
	var ok bool

	value, ok = myCache.Load("foo")
	fmt.Printf("%d: %v\n", value, ok)

	value, ok = myCache.Load("bar")
	fmt.Printf("%d: %v\n", value, ok)

	value, ok = myCache.Load("foobar")
	fmt.Printf("%d: %v\n", value, ok)

	// Output:
	// 13: true
	// 42: true
	// 0: false
}

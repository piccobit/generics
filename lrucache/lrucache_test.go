package lrucache_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/piccobit/generics/lrucache"
)

type car struct {
	name       string
	colour     string
	horsepower int
}

func (p *car) ID() string {
	return p.name
}

func ExampleLRUCache_Add_string() {
	var err error

	myStringLRU := lrucache.New[string](10)

	for i := 1; i <= 10; i++ {
		_, err = myStringLRU.Add(strconv.Itoa(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	_, err = myStringLRU.Add("11")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myStringLRU)
	// Output:
	// Content: [2,3,4,5,6,7,8,9,10,11]
}

func ExampleLRUCache_Add_string_2() {
	var err error

	myStringLRU := lrucache.New[string](10)

	id := make([]string, 10)

	for i := 1; i <= 10; i++ {
		id[i-1], err = myStringLRU.Add(strconv.Itoa(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	err = myStringLRU.AddByID(id[4], "5")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myStringLRU)
	// Output:
	// Content: [1,2,3,4,6,7,8,9,10,5]
}

func ExampleLRUCache_Add_string_3() {
	var err error

	myStringLRU := lrucache.New[string](10)

	id := make([]string, 10)

	for i := 1; i <= 10; i++ {
		id[i-1], err = myStringLRU.Add(strconv.Itoa(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	err = myStringLRU.AddByID(id[4], "5")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myStringLRU.AddByID(id[0], "1")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myStringLRU)
	// Output:
	// Content: [2,3,4,6,7,8,9,10,5,1]
}

func ExampleLRUCache_AddByID_string() {
	var err error

	myStringLRU := lrucache.New[string](10)

	for i := 1; i <= 10; i++ {
		err = myStringLRU.AddByID(strconv.Itoa(i), strconv.Itoa(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	err = myStringLRU.AddByID("11", "11")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myStringLRU)
	// Output:
	// Content: [2,3,4,5,6,7,8,9,10,11]
}

func ExampleLRUCache_AddByID_string_2() {
	var err error

	myStringLRU := lrucache.New[string](10)

	for i := 1; i <= 10; i++ {
		err = myStringLRU.AddByID(strconv.Itoa(i), strconv.Itoa(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	err = myStringLRU.AddByID("5", "5")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myStringLRU)
	// Output:
	// Content: [1,2,3,4,6,7,8,9,10,5]
}

func ExampleLRUCache_AddByID_car() {
	var err error

	myCarLRU := lrucache.New[car](10)

	err = myCarLRU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLRU.AddByID("Corvette", car{
		name:       "Little",
		colour:     "red",
		horsepower: 200,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	fmt.Printf("Content: %v", myCarLRU)
	// Output:
	// Content: [{Beetle blue 60},{Little red 200}]
}

func ExampleLRUCache_AddByID_car_2() {
	var err error

	myCarLRU := lrucache.New[car](10)

	err = myCarLRU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLRU.AddByID("Corvette", car{
		name:       "Little",
		colour:     "red",
		horsepower: 200,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLRU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	for k, v := range myCarLRU.GetCache() {
		fmt.Printf("%s: %v\n", k, v)
	}
	// Output:
	// Corvette: {Little red 200}
	// VW: {Beetle blue 60}
}

func ExampleLRUCache_Get_car() {
	var err error

	myCarLRU := lrucache.New[car](10)

	err = myCarLRU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLRU.AddByID("Corvette", car{
		name:       "Little",
		colour:     "red",
		horsepower: 200,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLRU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	_, _ = myCarLRU.Get("Corvette")

	for k, v := range myCarLRU.GetCache() {
		fmt.Printf("%s: %v\n", k, v)
	}
	// Output:
	// Corvette: {Little red 200}
}

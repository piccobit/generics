package lfucache_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/piccobit/generics/lfucache"
)

type car struct {
	name       string
	colour     string
	horsepower int
}

func (p *car) ID() string {
	return p.name
}

func ExampleLFUCache_Add_string() {
	var err error

	myStringLFU := lfucache.New[string](10)

	ids := make(map[string]string)

	for i := 1; i <= 10; i++ {
		id := strconv.Itoa(i)
		ids[id], err = myStringLFU.Add(id)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	for i := 1; i <= 10; i++ {
		if i == 4 {
			continue
		}

		id := strconv.Itoa(i)
		_, ok := myStringLFU.Get(ids[id])
		if !ok {
			fmt.Println(fmt.Errorf("could not find id '%s'", id))
		}
	}

	_, err = myStringLFU.Add("11")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	for _, v := range myStringLFU.GetCache() {
		fmt.Printf("%v\n", v)
	}

	// Unordered output:
	// 1
	// 2
	// 3
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
}

func ExampleLFUCache_AddByID_string() {
	var err error

	myStringLFU := lfucache.New[string](10)

	for i := 1; i <= 10; i++ {
		id := strconv.Itoa(i)
		err = myStringLFU.AddByID(id, id)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
		}
	}

	for i := 1; i <= 10; i++ {
		if i == 4 {
			continue
		}

		id := strconv.Itoa(i)
		_, ok := myStringLFU.Get(id)
		if !ok {
			fmt.Println(fmt.Errorf("could not find id '%s'", id))
		}
	}

	err = myStringLFU.AddByID("11", "11")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	for _, v := range myStringLFU.GetCache() {
		fmt.Printf("%v\n", v)
	}

	// Unordered output:
	// 1
	// 2
	// 3
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
}

func ExampleLFUCache_AddByID_car() {
	var err error

	myCarLFU := lfucache.New[car](3)

	err = myCarLFU.AddByID("VW", car{
		name:       "Beetle",
		colour:     "blue",
		horsepower: 60,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLFU.AddByID("Corvette", car{
		name:       "Little",
		colour:     "red",
		horsepower: 200,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	err = myCarLFU.AddByID("Aston Martin", car{
		name:       "James Bond",
		colour:     "black",
		horsepower: 150,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	_, _ = myCarLFU.Get("VW")
	_, _ = myCarLFU.Get("Aston Martin")

	err = myCarLFU.AddByID("Rolls Royce", car{
		name:       "Lisbeth",
		colour:     "silver",
		horsepower: 42,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\\n", err.Error())
	}

	for _, v := range myCarLFU.GetCache() {
		fmt.Printf("%v\n", v)
	}

	// Unordered output:
	// {James Bond black 150}
	// {Beetle blue 60}
	// {Lisbeth silver 42}
}

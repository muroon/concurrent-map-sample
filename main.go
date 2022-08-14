package main

import (
	"fmt"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	concurrentMap()
	syncMap()
}

func syncMap() {
	m := sync.Map{}

	m.Store("foo", "bar")

	bar, ok := m.Load("foo")
	if !ok {
		fmt.Println("[syncMap] cannot Get foo")
	}

	fmt.Printf("[syncMap] %T, %v\n", bar, bar)

	m.Delete("foo")
	bar, ok = m.Load("foo")
	if !ok {
		fmt.Println("[syncMap] cannot Get foo")
	}

	fmt.Printf("[syncMap] %T, %v\n", bar, bar)
}

func concurrentMap() {
	// Create a new map.
	m := cmap.New[string]()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")

	// Retrieve item from map.
	bar, ok := m.Get("foo")
	if !ok {
		fmt.Println("[concurrentMap] cannot Get foo")
	}

	fmt.Printf("[concurrentMap] %T, %v\n", bar, bar)

	m.Remove("foo")
	bar, ok = m.Get("foo")
	if !ok {
		fmt.Println("[concurrentMap] cannot Get foo")
	}

	fmt.Printf("[concurrentMap] %T, %v\n", bar, bar)
}

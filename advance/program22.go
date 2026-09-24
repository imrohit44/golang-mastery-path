package main

import (
	"fmt"
	"sync"
)

func main() {
	var cache sync.Map
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			cache.Store(fmt.Sprintf("key_%d", val), val*10)
		}(i)
	}
	wg.Wait()

	// Load or Store (Atomic operation)
	actual, loaded := cache.LoadOrStore("key_0", 999)
	fmt.Printf("Key_0 value: %v (Was it already there? %t)\n", actual, loaded)

	// Iterate over the concurrent map safely
	cache.Range(func(key, value interface{}) bool {
		fmt.Printf("Cache Entry -> %s: %v\n", key, value)
		return true // Continue iteration
	})
}
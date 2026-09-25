package main

import "fmt"

func main() {
	// Make a slice with length 0, but capacity 5
	data := make([]int, 0, 5)

	fmt.Printf("Initial - Length: %d, Capacity: %d\n", len(data), cap(data))

	// Appending items up to the capacity
	for i := 1; i <= 5; i++ {
		data = append(data, i)
	}
	fmt.Printf("Filled  - Length: %d, Capacity: %d\n", len(data), cap(data))

	// Appending beyond capacity forces Go to allocate a new, larger backing array
	data = append(data, 6)
	fmt.Printf("Grown   - Length: %d, Capacity: %d\n", len(data), cap(data))
}
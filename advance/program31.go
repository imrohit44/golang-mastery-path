package main

import "fmt"

// Map applies a function to every item in a slice and returns a new slice
func Map[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))
	for i, v := range items {
		result[i] = transform(v)
	}
	return result
}

// Filter returns a slice containing only elements that match the predicate
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Filter even numbers
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	
	// Map to their string representations
	strings := Map(evens, func(n int) string { return fmt.Sprintf("Number: %d", n) })

	fmt.Println("Filtered and Mapped:", strings)
}
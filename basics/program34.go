package main

import "fmt"

func main() {
	source := []int{1, 2, 3}
	
	// Destination slice must be allocated with enough length
	destination := make([]int, len(source))
	
	copiedCount := copy(destination, source)

	fmt.Printf("Copied %d elements.\n", copiedCount)
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)

	// Modifying source does NOT affect destination
	source[0] = 99
	fmt.Println("After modifying source[0]:")
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}
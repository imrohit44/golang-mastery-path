package main

import "fmt"

func main() {
	// Anonymous function executed immediately
	func(msg string) {
		fmt.Println("Immediate:", msg)
	}("Hello there!")

	// Anonymous function assigned to a variable (Closure)
	counter := 0
	increment := func() int {
		counter++ // Captures the 'counter' variable from the outer scope
		return counter
	}

	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
	fmt.Println("Count:", increment())
}
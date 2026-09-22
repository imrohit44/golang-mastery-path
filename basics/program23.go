package main

import "fmt"

func safeDivide(a, b int) {
	// Defer a function to recover from any panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	if b == 0 {
		panic("Cannot divide by zero!")
	}
	
	fmt.Println("Result:", a/b)
}

func main() {
	safeDivide(10, 2)
	safeDivide(10, 0) // This will panic, but the program will recover and continue
	fmt.Println("Program finished successfully.")
}
package main

import "fmt"

// Type Definition: Creates a distinct new type
type CustomInt int

// Type Alias: Creates an exact synonym for an existing type
type IntegerAlias = int

func main() {
	var a int = 10
	var b CustomInt = 20
	var c IntegerAlias = 30

	// a = b // ERROR: Cannot use CustomInt as int in assignment
	a = int(b) // Must explicitly convert

	a = c // Works perfectly! IntegerAlias is just another name for int

	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
}
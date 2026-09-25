package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Outer x before block:", x)

	if true {
		x := 50 // This creates a NEW 'x' inside this block, shadowing the outer 'x'
		fmt.Println("Inner x inside block:", x)
	}

	fmt.Println("Outer x after block:", x) // Still 10!
}
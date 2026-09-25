package main

import "fmt"

func main() {
	inventory := map[string]int{
		"apples":  5,
		"oranges": 10,
	}

	// Check for a key that exists
	val, ok := inventory["apples"]
	if ok {
		fmt.Printf("Apples found: %d\n", val)
	}

	// Check for a key that does NOT exist
	val, ok = inventory["bananas"]
	if !ok {
		fmt.Println("Bananas are not in the inventory.")
	}
}
package main

import "fmt"

func analyzeValue(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Printf("It's an integer multiplied by 2: %d\n", v*2)
	case string:
		fmt.Printf("It's a string of length %d: %s\n", len(v), v)
	case bool:
		fmt.Printf("It's a boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	analyzeValue(42)
	analyzeValue("Golang")
	analyzeValue(true)
	analyzeValue(3.14)
}
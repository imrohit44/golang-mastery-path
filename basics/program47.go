package main

import "fmt"

// Accepts any type of data
func printAnything(data interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", data, data)
}

func main() {
	printAnything(42)
	printAnything("Golang")
	printAnything(3.14159)
	printAnything([]int{1, 2, 3})
}
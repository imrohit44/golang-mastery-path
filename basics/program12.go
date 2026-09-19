package main

import "fmt"

func main() {
	defer fmt.Println("3. This runs just before the function exits.")
	
	fmt.Println("1. This runs first.")
	fmt.Println("2. This runs second.")
}
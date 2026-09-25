package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50}

	fmt.Println("Original:", numbers)
	fmt.Println("Index 1 to 4:", numbers[1:4]) // Includes index 1, 2, 3
	fmt.Println("Start to index 3:", numbers[:3]) // Equivalent to [0:3]
	fmt.Println("Index 3 to end:", numbers[3:]) // Equivalent to [3:len(numbers)]
}
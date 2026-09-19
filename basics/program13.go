package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum of 1, 2, 3:", sumNumbers(1, 2, 3))
	fmt.Println("Sum of 10, 20:", sumNumbers(10, 20))
}
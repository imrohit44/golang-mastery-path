package main

import "fmt"

func main() {
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("Breaking entirely out of the nested loops!")
				break OuterLoop // Breaks the outer loop, not just the inner one
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	fmt.Println("Loop execution finished.")
}
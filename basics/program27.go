package main

import "fmt"

func main() {
	// Create a 2D slice (a matrix)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Value at row 1, col 2:", matrix[1][2]) // 6

	// Iterate through the matrix
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("Matrix[%d][%d] = %d\n", i, j, val)
		}
	}
}
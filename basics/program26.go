package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Pi:", math.Pi)
	
	fmt.Println("Square root of 16:", math.Sqrt(16))
	
	fmt.Println("2 to the power of 3:", math.Pow(2, 3))
	
	fmt.Println("Absolute value of -42:", math.Abs(-42))
	
	fmt.Println("Max of 10 and 20:", math.Max(10, 20))
	
	// Rounding
	fmt.Println("Round 3.6:", math.Round(3.6))
}
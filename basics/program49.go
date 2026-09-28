package main

import "fmt"

func main() {
	// Create complex numbers using the complex() function
	c1 := complex(5, 7) // 5 + 7i
	c2 := complex(3, 2) // 3 + 2i

	// You can also use the 'i' syntax directly
	c3 := 2 + 4i

	fmt.Println("c1:", c1)
	fmt.Println("c1 + c2 =", c1+c2)
	fmt.Println("c1 * c3 =", c1*c3)

	// Extract real and imaginary parts
	fmt.Printf("c1 Real: %f, Imaginary: %f\n", real(c1), imag(c1))
}
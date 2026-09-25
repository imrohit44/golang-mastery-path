package main

import "fmt"

func main() {
	a := 12 // Binary: 1100
	b := 10 // Binary: 1010

	fmt.Printf("a & b  (AND): %04b\n", a&b)   // 1000
	fmt.Printf("a | b   (OR): %04b\n", a|b)   // 1110
	fmt.Printf("a ^ b  (XOR): %04b\n", a^b)   // 0110
	fmt.Printf("a << 1 (Left Shift): %04b\n", a<<1) // 11000
	fmt.Printf("a >> 1 (Right Shift): %04b\n", a>>1) // 0110
}
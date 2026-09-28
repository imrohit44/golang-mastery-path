package main

import "fmt"

func main() {
	// Standard string literal (requires escape characters)
	normalString := "Line 1\nLine 2\n\"Quoted Text\""

	// Raw string literal (preserves exact formatting)
	rawString := `
Line 1
Line 2
"Quoted Text"
\n \t (these are literal, not escapes)`

	fmt.Println("Normal String:")
	fmt.Println(normalString)
	
	fmt.Println("\nRaw String:")
	fmt.Println(rawString)
}
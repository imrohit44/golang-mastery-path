package main

import "fmt"

func main() {
	// This string contains English letters and a Japanese kanji (multi-byte)
	text := "Go 世界"

	fmt.Println("Byte length:", len(text)) // Counts bytes, not characters

	// Ranging over a string automatically decodes it into Runes
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", index, char, char)
	}
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Go", "is", "an", "efficient", "language."}
	
	var builder strings.Builder

	for i, word := range words {
		builder.WriteString(word)
		if i < len(words)-1 {
			builder.WriteString(" ")
		}
	}

	finalString := builder.String()
	fmt.Println("Built string:", finalString)
}
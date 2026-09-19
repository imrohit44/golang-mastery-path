package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go programming is fun"

	fmt.Println("Uppercase:", strings.ToUpper(text))
	fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))
	fmt.Println("Replace 'fun' with 'awesome':", strings.ReplaceAll(text, "fun", "awesome"))
}
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"black": "#000000",
	}

	fmt.Println("First Iteration:")
	for name, hex := range colors {
		fmt.Printf("%s: %s\n", name, hex)
	}

	fmt.Println("\nSecond Iteration (Order may change!):")
	for name, hex := range colors {
		fmt.Printf("%s: %s\n", name, hex)
	}
}
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is always the path to the program itself
	args := os.Args

	fmt.Println("Total arguments:", len(args))
	fmt.Println("Program name:", args[0])

	if len(args) > 1 {
		fmt.Println("First user argument:", args[1])
		fmt.Println("All user arguments:", args[1:])
	} else {
		fmt.Println("No user arguments provided. Try running: go run main.go hello world")
	}
}
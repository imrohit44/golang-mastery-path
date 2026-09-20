package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	namePtr := flag.String("name", "Guest", "A name to say hello to")
	agePtr := flag.Int("age", 0, "Your age")
	verbosePtr := flag.Bool("verbose", false, "Enable verbose output")

	// Parse flags
	flag.Parse()

	fmt.Printf("Hello, %s! (Age: %d)\n", *namePtr, *agePtr)

	if *verbosePtr {
		fmt.Println("Verbose mode is enabled.")
	}
}
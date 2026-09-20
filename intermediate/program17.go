package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 2
	ch := make(chan string, 2)

	// These won't block because there is room in the buffer
	ch <- "Message 1"
	ch <- "Message 2"

	// Read from the channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
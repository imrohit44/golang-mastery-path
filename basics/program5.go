package main

import "fmt"

func main() {
    // Slices are dynamic arrays in Go
    fruits := []string{"Apple", "Banana", "Cherry"}
    fruits = append(fruits, "Orange")

    for index, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", index, fruit)
    }
}
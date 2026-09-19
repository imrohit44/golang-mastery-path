package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the generator using the current time
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random number (0-99):", rand.Intn(100))
	fmt.Println("Random float (0.0-1.0):", rand.Float64())
}
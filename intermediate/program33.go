package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "super_secret_password"

	// Create a new hash instance
	hash := sha256.New()
	
	// Write data to it (must be bytes)
	hash.Write([]byte(password))
	
	// Compute the final hash (nil means don't append to an existing byte slice)
	hashBytes := hash.Sum(nil)

	// Format as a hexadecimal string
	hashString := fmt.Sprintf("%x", hashBytes)
	
	fmt.Println("Original:", password)
	fmt.Println("SHA-256:", hashString)
}
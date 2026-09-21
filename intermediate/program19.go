package main

import (
	"fmt"
	"os"
)

func main() {
	// Set an environment variable
	os.Setenv("APP_PORT", "8080")

	// Get an environment variable
	port := os.Getenv("APP_PORT")
	fmt.Println("Server configured to run on port:", port)

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Fallback default
	}
	fmt.Println("Database host:", dbHost)
}
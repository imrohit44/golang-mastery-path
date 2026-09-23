package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the regex once (MustCompile panics if the syntax is invalid)
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	emails := []string{
		"valid.user@example.com",
		"invalid-email@com",
		"hello@world.org",
	}

	for _, email := range emails {
		isValid := emailRegex.MatchString(email)
		fmt.Printf("%-25s : Valid? %t\n", email, isValid)
	}

	// Extracting data
	text := "Contact us at support@example.com or sales@example.com"
	matches := emailRegex.FindAllString(text, -1)
	fmt.Println("\nFound emails:", matches)
}
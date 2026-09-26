package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://api.example.com:8080/v1/users?role=admin&active=true#profile"

	parsed, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsed.Scheme)
	fmt.Println("Host:", parsed.Host)
	fmt.Println("Path:", parsed.Path)
	fmt.Println("Fragment:", parsed.Fragment)

	// Extracting Query Parameters
	queries := parsed.Query()
	fmt.Println("Role Query:", queries.Get("role"))
	fmt.Println("Active Query:", queries.Get("active"))
}
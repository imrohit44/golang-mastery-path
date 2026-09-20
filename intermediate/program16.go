package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Make a GET request
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response starts with: %s...\n", string(body[:50]))
}
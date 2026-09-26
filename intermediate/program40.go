package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	// Create a new request object instead of using http.Get directly
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		panic(err)
	}

	// Add custom headers
	req.Header.Add("Authorization", "Bearer my-secret-token")
	req.Header.Add("User-Agent", "Go-Mastery-Bot/1.0")
	req.Header.Add("X-Custom-Header", "GolangRocks")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Server received the following headers:")
	fmt.Println(string(body))
}
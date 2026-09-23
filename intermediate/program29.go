package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Configure an HTTP client with a strict 2-second timeout
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	// This site delays the response by 3 seconds, which will trigger our timeout
	url := "https://httpstat.us/200?sleep=3000"
	
	fmt.Println("Making request to", url)
	resp, err := client.Get(url)
	
	if err != nil {
		fmt.Println("Request failed (likely timed out):")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)
}
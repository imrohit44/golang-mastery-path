package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetWithRetry(url string, maxRetries int) (*http.Response, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	var err error
	var resp *http.Response

	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Attempt %d making request to %s\n", attempt, url)
		resp, err = client.Get(url)
		
		if err == nil && resp.StatusCode == 200 {
			return resp, nil // Success
		}

		fmt.Printf("Attempt %d failed. Retrying...\n", attempt)
		time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff
	}

	return nil, fmt.Errorf("failed after %d retries: %v", maxRetries, err)
}

func main() {
	// A deliberately non-existent domain to trigger retries
	url := "http://this-domain-surely-does-not-exist.org"
	
	_, err := GetWithRetry(url, 3)
	if err != nil {
		fmt.Println("Final Error:", err)
	}
}
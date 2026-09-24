package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "https://github.com", nil)

	var start time.Time
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Printf("DNS Lookup for %s started...\n", info.Host)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Lookup finished in %v\n", time.Since(start))
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Connection established in %v\n", time.Since(start))
		},
	}

	// Attach the trace to the request context
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start = time.Now()
	client := http.Client{Timeout: 5 * time.Second}
	_, err := client.Do(req)
	
	if err != nil {
		fmt.Println("Request failed:", err)
	}
}
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// The target server we are proxying to
	targetURL, _ := url.Parse("https://google.com")
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Modify the request before forwarding it
	proxy.Director = func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Proxying request: %s\n", r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Reverse proxy listening on :8080. Traffic forwarded to google.com")
	// http.ListenAndServe(":8080", nil)
}
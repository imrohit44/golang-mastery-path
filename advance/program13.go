package main

import (
	"fmt"
	"time"
)

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

// Option is a function that modifies a Server
type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) { s.Port = port }
}

func WithTimeout(t time.Duration) Option {
	return func(s *Server) { s.Timeout = t }
}

func NewServer(host string, options ...Option) *Server {
	// Set defaults
	srv := &Server{
		Host:    host,
		Port:    8080,
		Timeout: 30 * time.Second,
	}
	
	// Apply overrides
	for _, opt := range options {
		opt(srv)
	}
	return srv
}

func main() {
	srv := NewServer("localhost", WithPort(9000), WithTimeout(5*time.Second))
	fmt.Printf("Configured Server: %+v\n", srv)
}
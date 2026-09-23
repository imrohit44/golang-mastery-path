// Normally placed in math_test.go
package main

import (
	"fmt"
	"testing"
)

func IsEven(n int) bool {
	return n%2 == 0
}

func TestIsEven(t *testing.T) {
	// Table of test cases
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"positive even", 2, true},
		{"positive odd", 3, false},
		{"zero", 0, true},
		{"negative odd", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func main() {
	fmt.Println("This is a test file. Run using: go test")
}
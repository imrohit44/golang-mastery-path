// Normally, this would be in math_test.go
package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// Test function must start with "Test" and take a *testing.T pointer
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func main() {
	// Run "go test" in the terminal to execute tests.
	fmt.Println("This file contains unit tests. Run `go test`.")
}
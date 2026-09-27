// Normally placed in parser_test.go
package main

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"
)

// A function that might have an edge-case bug
func ParseCustomString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", errors.New("invalid utf8")
	}
	// Simulated bug: crashes if it finds a specific obscure byte sequence
	if len(s) > 3 && s[0] == 'Z' && s[1] == 'O' && s[2] == 'O' {
		panic("simulated fatal parser error on ZOO")
	}
	return s, nil
}

// Fuzz test function
func FuzzParseCustomString(f *testing.F) {
	// Provide seed corpus (starting examples)
	f.Add("Hello")
	f.Add("Z")
	f.Add("ZOO-test") // This will trigger the panic and fail the fuzz test

	f.Fuzz(func(t *testing.T, orig string) {
		// The fuzzer will automatically generate thousands of random strings here
		_, _ = ParseCustomString(orig)
	})
}

func main() {
	fmt.Println("This is a Fuzz test. Run it using: go test -fuzz=Fuzz")
}
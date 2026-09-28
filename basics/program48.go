package main

import "fmt"

const (
	TypedConst   int = 100 // Strictly an int
	UntypedConst     = 100 // Can adapt to int, int64, float64, etc.
)

func main() {
	var a int32 = TypedConst // ERROR if TypedConst is int, but we can't easily assign it.
	_ = a // (Ignoring error for demo context)

	// Untyped flexibly adapts to int32, float64, etc.
	var b int32 = UntypedConst
	var c float64 = UntypedConst

	fmt.Printf("b (int32): %d\n", b)
	fmt.Printf("c (float64): %.2f\n", c)
}
package main

import (
	"fmt"
	"reflect"
	"time"
)

// MakeTimed takes a pointer to a function variable and dynamically rewrites 
// it to wrap the original logic with a timer.
func MakeTimed(fn interface{}) {
	fnValue := reflect.ValueOf(fn).Elem()
	originalFn := fnValue.Interface()

	// Create a new function dynamically
	wrapper := reflect.MakeFunc(fnValue.Type(), func(args []reflect.Value) []reflect.Value {
		start := time.Now()
		
		// Call the original function
		results := reflect.ValueOf(originalFn).Call(args)
		
		fmt.Printf("[Timer] Execution took %v\n", time.Since(start))
		return results
	})

	// Replace the function pointer with our new dynamically generated function
	fnValue.Set(wrapper)
}

func main() {
	// A standard function variable
	var complexMath func(int, int) int

	// Assign the actual logic
	complexMath = func(a, b int) int {
		time.Sleep(100 * time.Millisecond) // Simulate heavy work
		return a * b
	}

	// Dynamically wrap it with our timer!
	MakeTimed(&complexMath)

	result := complexMath(5, 10)
	fmt.Println("Result:", result)
}
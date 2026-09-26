package main

import "fmt"

func calculateAndAudit(a, b int) (result int) {
	// This defer runs after the return statement is evaluated, 
	// but before the value is actually handed back to the caller.
	defer func() {
		fmt.Println("Audit Log: Calculation performed, overriding result for demo.")
		result = result * 10 // Modifying the named return value
	}()

	return a + b // Normally returns 5
}

func main() {
	final := calculateAndAudit(2, 3)
	fmt.Println("Final returned value:", final)
}
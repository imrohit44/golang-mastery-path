package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to Integer
	strAge := "25"
	age, err := strconv.Atoi(strAge)
	if err == nil {
		fmt.Printf("Parsed Age: %d (Type: %T)\n", age, age)
	}

	// Integer to String
	year := 2026
	strYear := strconv.Itoa(year)
	fmt.Printf("String Year: %s (Type: %T)\n", strYear, strYear)

	// String to Boolean
	strBool := "true"
	b, _ := strconv.ParseBool(strBool)
	fmt.Printf("Parsed Boolean: %t (Type: %T)\n", b, b)
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi returns (int, error). We only want the int and ignore the error.
	number, _ := strconv.Atoi("100")
	
	fmt.Println("Parsed number:", number)
	
	// Ignoring the index in a range loop
	names := []string{"Alice", "Bob"}
	for _, name := range names {
		fmt.Println("Name:", name)
	}
}
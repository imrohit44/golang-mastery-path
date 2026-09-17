package main

import (
	"fmt"
	"unsafe"
)

func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	// Direct unsafely pointed conversion without copying underlying array
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func main() {
	str := "High Performance Go"
	bytes := StringToBytes(str)

	fmt.Printf("String: %s\n", str)
	fmt.Printf("Byte Array: %v\n", bytes)
}
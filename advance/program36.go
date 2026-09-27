package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [4]int{10, 20, 30, 40}

	// Get a raw pointer to the first element
	basePointer := unsafe.Pointer(&array[0])
	
	// Size of an integer in bytes (8 bytes on a 64-bit system)
	intSize := unsafe.Sizeof(array[0])

	fmt.Println("Iterating array using raw memory pointer arithmetic:")
	for i := 0; i < len(array); i++ {
		// Calculate the physical memory address of the next element
		// uintptr is an integer type large enough to hold the bit pattern of any pointer
		nextAddress := uintptr(basePointer) + uintptr(i)*intSize
		
		// Convert the calculated address back to an unsafe.Pointer, then to a typed *int pointer
		valPointer := (*int)(unsafe.Pointer(nextAddress))
		
		fmt.Printf("Index %d is stored at %v with value %d\n", i, valPointer, *valPointer)
	}
}
package main

import (
	"fmt"
	"sync"
)

// Create a pool of 4KB byte slices
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 4096)
		return &b
	},
}

func ProcessStreamData(data string) {
	// Borrow a byte slice from the pool
	bufPtr := bytePool.Get().(*[]byte)
	buf := *bufPtr
	
	// Ensure we return the slice to the pool when the function exits
	defer bytePool.Put(bufPtr)

	// Use the buffer (simulated copy)
	length := copy(buf, data)
	fmt.Printf("Processed %d bytes using pooled memory\n", length)
	
	// Reset the slice before returning it to prevent data leaks
	for i := 0; i < length; i++ {
		buf[i] = 0
	}
}

func main() {
	ProcessStreamData("High performance I/O stream 1")
	ProcessStreamData("High performance I/O stream 2")
}
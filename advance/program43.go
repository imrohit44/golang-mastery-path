package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Create a dummy file
	filename := "mmap_example.dat"
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	file.Write([]byte("Hello, Memory Mapped World! This is ultra-fast."))
	defer os.Remove(filename)

	// Get file stats to know the exact size to map
	stat, _ := file.Stat()
	size := stat.Size()

	// Map the file into memory
	// PROT_READ|PROT_WRITE means we can read and write to this memory segment
	// MAP_SHARED means changes propagate back to the underlying file
	data, err := syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original Memory: %s\n", string(data))

	// Modify the RAM directly. This automatically writes to the file on disk!
	data[0] = 'Y'
	data[1] = 'e'
	data[2] = 'l'
	data[3] = 'l'
	data[4] = 'o'

	fmt.Printf("Modified Memory: %s\n", string(data))

	// Clean up and unmap the memory
	syscall.Munmap(data)
	file.Close()
}
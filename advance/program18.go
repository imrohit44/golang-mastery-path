package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	printMemStats()

	// Simulate heavy allocation
	data := make([][]int, 0)
	for i := 0; i < 10000; i++ {
		data = append(data, make([]int, 10000))
	}
	
	time.Sleep(time.Second) // Let system settle
	fmt.Println("After Allocation:")
	printMemStats()
	
	// Force Garbage Collection
	runtime.GC()
	fmt.Println("After Garbage Collection:")
	printMemStats()
}
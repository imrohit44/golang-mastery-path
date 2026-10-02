package main

import (
	"fmt"
	"runtime"
)

func PrintStackTrace() {
	// Allocate a slice to hold the program counters (memory addresses of function calls)
	pcs := make([]uintptr, 10)
	
	// Skip 0 (Callers itself) and 1 (PrintStackTrace) to get the caller's frames
	depth := runtime.Callers(2, pcs)
	
	fmt.Println("--- Stack Trace ---")
	
	frames := runtime.CallersFrames(pcs[:depth])
	for {
		frame, more := frames.Next()
		fmt.Printf("Func: %s\n\tFile: %s:%d\n", frame.Function, frame.File, frame.Line)
		
		if !more {
			break
		}
	}
	fmt.Println("-------------------")
}

func Alpha() { Beta() }
func Beta()  { Gamma() }
func Gamma() { PrintStackTrace() }

func main() {
	Alpha()
}
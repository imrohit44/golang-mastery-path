package main

import (
	"fmt"
	"sort"
)

type Server struct {
	Name string
	Load int
}

// Create a custom slice type
type ByLoad []Server

// Implement sort.Interface
func (a ByLoad) Len() int           { return len(a) }
func (a ByLoad) Less(i, j int) bool { return a[i].Load < a[j].Load }
func (a ByLoad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	servers := []Server{
		{"Server Alpha", 85},
		{"Server Beta", 12},
		{"Server Gamma", 45},
	}

	fmt.Println("Before Sort:", servers)

	// Sort using our custom logic
	sort.Sort(ByLoad(servers))

	fmt.Println("After Sort (Ascending Load):", servers)
}
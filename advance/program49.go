package main

import (
	"container/heap"
	"fmt"
)

// Item represents a scheduled task
type Item struct {
	Value    string
	Priority int
	index    int // Used internally by the heap package
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest priority, so we use >
	return pq[i].Priority > pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Insert items in random order
	heap.Push(&pq, &Item{Value: "Low priority task", Priority: 1})
	heap.Push(&pq, &Item{Value: "CRITICAL SYSTEM FAILURE", Priority: 99})
	heap.Push(&pq, &Item{Value: "Medium priority task", Priority: 5})

	// Pop them out. They will come out ordered by priority!
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Processing Priority %d: %s\n", item.Priority, item.Value)
	}
}
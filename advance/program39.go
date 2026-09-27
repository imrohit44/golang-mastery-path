package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Iterator returns a stateful function that yields the next value and a boolean
func (ll *LinkedList) Iterator() func() (int, bool) {
	current := ll.Head
	return func() (int, bool) {
		if current == nil {
			return 0, false // Reached the end
		}
		val := current.Value
		current = current.Next // Advance state
		return val, true
	}
}

func main() {
	list := &LinkedList{
		Head: &Node{Value: 10, Next: &Node{Value: 20, Next: &Node{Value: 30, Next: nil}}},
	}

	// Create the iterator
	next := list.Iterator()

	// Consume the iterator
	for {
		val, ok := next()
		if !ok {
			break
		}
		fmt.Println("List item:", val)
	}
}
package main

import (
	"fmt"
)

// Stack can hold any type T
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // Returns the zero value for type T
		return zero, false
	}
	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[:index]
	return item, true
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	
	val, _ := intStack.Pop()
	fmt.Println("Popped Integer:", val)

	strStack := Stack[string]{}
	strStack.Push("Hello")
	strStack.Push("Generics")
	
	valStr, _ := strStack.Pop()
	fmt.Println("Popped String:", valStr)
}
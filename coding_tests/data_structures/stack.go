package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	// Create a new stack
	s := stack.New()

	// Push elements onto the stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Peek at the top element on the stack
	top := s.Peek()
	fmt.Println("Top element:", top)

	// Pop elements from the stack
	for s.Len() > 0 {
		val := s.Pop()
		fmt.Println(val)
	}
}

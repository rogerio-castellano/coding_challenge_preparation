package main

import (
	"fmt"
)

// Stack represents a basic stack.
type Stack []string

// Push adds an element to the top of the stack.
func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return "" // Stack is empty
	}
	n := len(*s) - 1
	top := (*s)[n]
	*s = (*s)[:n]
	return top
}

func main() {
	var stack Stack

	// Push elements onto the stack
	stack.Push("world!")
	stack.Push("wonderful")
	stack.Push("Hello")

	// Pop elements from the stack until it's empty
	for len(stack) > 0 {
		fmt.Println(stack.Pop())
	}
	// Output: Hello world!
}

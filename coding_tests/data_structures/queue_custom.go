package main

import (
	"fmt"
)

// Queue represents a basic queue.
type Queue []string

// Enqueue adds an element to the back of the queue.
func (q *Queue) Enqueue(item string) {
	*q = append(*q, item)
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue) Dequeue() string {
	if len(*q) == 0 {
		return "" // Queue is empty
	}
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func main() {
	var queue Queue

	// Enqueue elements into the queue
	queue.Enqueue("Hello")
	queue.Enqueue("wonderful")
	queue.Enqueue("world!")

	// Dequeue elements from the queue until it's empty
	for len(queue) > 0 {
		fmt.Println(queue.Dequeue())
	}
	// Output: Helloworld!
}

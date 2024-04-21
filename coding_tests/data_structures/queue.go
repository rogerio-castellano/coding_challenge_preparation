package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

func main() {
	// Create a new queue
	q := queue.New()

	// Enqueue elements into the queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue elements from the queue
	for q.Len() > 0 {
		val := q.Dequeue()
		fmt.Println(val)
	}
}

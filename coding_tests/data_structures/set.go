package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

func main() {
	// Create a new set
	s := set.New()

	// Insert elements into the set
	s.Insert("apple")
	s.Insert("banana")
	s.Insert("cherry")

	// Check if elements are in the set
	fmt.Println("Does the set contain 'banana'?", s.Has("banana"))

	// Remove an element from the set
	s.Remove("apple")

	// Iterate over the set
	s.Do(func(value interface{}) {
		fmt.Println(value)
	})

	// Create another set to demonstrate set operations
	t := set.New()
	t.Insert("banana")
	t.Insert("date")

	// Intersection of sets
	u := s.Intersection(t)
	fmt.Println("Intersection:", u)

	// Union of sets
	v := s.Union(t)
	fmt.Println("Union:", v)

	// Difference of sets
	w := s.Difference(t)
	fmt.Println("Difference:", w)
}

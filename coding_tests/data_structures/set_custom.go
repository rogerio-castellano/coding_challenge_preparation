package main

import (
	"fmt"
)

// Set represents a basic set.
type Set map[string]struct{}

// Add adds an element to the set.
func (s Set) Add(item string) {
	s[item] = struct{}{}
}

// Remove removes an element from the set.
func (s Set) Remove(item string) {
	delete(s, item)
}

// Contains checks if an element exists in the set.
func (s Set) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

func main() {
	mySet := make(Set)

	// Add elements to the set
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("cherry")

	// Check if an element exists
	fmt.Println("Contains 'banana':", mySet.Contains("banana"))

	// Remove an element
	mySet.Remove("cherry")

	// Print all elements
	for item := range mySet {
		fmt.Println(item)
	}
}

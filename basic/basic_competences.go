package main

import (
	"fmt"
	"time"
)

// Define a custom type and a struct
type Numeric int
type User struct {
	Name string
	Age  int
}

// Define an interface
type Greeter interface {
	Greet() string
}

// Implement the interface for the User struct
func (u User) Greet() string {
	return fmt.Sprintf("Hi, my name is %s and I am %d years old.", u.Name, u.Age)
}

func main() {
	// Common Golang Features
	fmt.Println("Exploring Go's features")

	// Understanding Basic Data structures
	var message string = "Learning Go!"
	var pointer *string
	var currentTime time.Time = time.Now()
	var number Numeric = 42
	slice := []string{"a", "b", "c"}

	fmt.Println(message, pointer, currentTime, number, slice)

	// Work with Collections
	array := [3]int{1, 2, 3}
	user := User{Name: "Alice", Age: 30}
	mapping := map[string]int{"one": 1, "two": 2}

	fmt.Println(array, user, mapping)

	// Passing values to functions
	displaySlice(slice)
	displayMap(mapping)

	// Types and Interfaces
	var greeter Greeter = user
	fmt.Println(greeter.Greet())
}

// Function to display a slice
func displaySlice(s []string) {
	fmt.Println("Slice contents:", s)
}

// Function to display a map
func displayMap(m map[string]int) {
	fmt.Println("Map contents:")
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

package main

import (
	"fmt"
)

func main() {
	fnstruct()
}

func fnstruct() {
	type person struct {
		name string
		age  int
	}

	person1 := person{name: "Rogerio", age: 52}

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(person1, dog)
}
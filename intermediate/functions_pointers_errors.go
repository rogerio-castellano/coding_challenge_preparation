package main

import (
	"errors"
	"fmt"
)

// Function with named return parameters and multiple return values
func divide(dividend, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result = dividend / divisor
	return
}

// Function demonstrating pointers
func increment(number *int) {
	*number++ // Increment the value at the pointer address
}

// Function demonstrating error handling
func safeDivide(dividend, divisor int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in safeDivide:", r)
		}
	}()

	result, err := divide(dividend, divisor)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Function Execution
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Working with Pointers
	num := 5
	increment(&num)
	fmt.Println("Incremented number:", num)

	// Error Handling
	fmt.Println("Safe Division Result:", safeDivide(10, 0))
}

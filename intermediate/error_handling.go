package main

import (
	"errors"
	"fmt"
	"log"
)

const conversion = 1.609

func main() {
	quotient, remainder, err := divide(10, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quotient, remainder)

	quotient, remainder, err = divide(10, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func divide(dividend int, divisor int) (quotient int, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("Divisor zero is invalid")
		return
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	err = nil
	return
}

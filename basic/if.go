package main

import (
	"fmt"
)

func main() {
	fnif()
}

func fnif() { //9 has 1 digit
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
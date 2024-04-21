package main

import (
	"fmt"
)

func main() {
	functionvariable()
}

func functionvariable() {
	var f func(int) int
	f = func(i int) int {
		if i == 0 {
			return 1
		}
		return i * f(i-1)
	}
	fmt.Println(f(2))
}

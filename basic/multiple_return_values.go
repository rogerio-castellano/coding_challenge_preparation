package main

import (
	"fmt"
)

func main() {
	fnmultiplereturnvalues()
}
func fnmultiplereturnvalues() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func vals() (int, int) {
	return 3, 7
}
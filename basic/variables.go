package main

import (
	"fmt"
)

func main() {
	fnvariables()
}

const sc string = "constant"

func fnvariables() {
	fmt.Println(sc)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d bool
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)
}

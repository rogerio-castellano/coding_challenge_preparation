package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Rogerio"
	result := ""
	for i := 0; i < len(text); i++ {
		if !(i%4 == 3) {
			result += (string(text[i]))
		} else {
			result += "_"
		}
	}
	fmt.Println(result) //Rog_rio
	result = ""

	for i, r := range []rune(text) {
		if !(i%4 == 3) {
			result += string(r)
		} else {
			result += "_"
		}
	}
	fmt.Println(result) //Rog_rio
	result = ""

	for i, char := range strings.Split(text, "") {
		if !(i%4 == 3) {
			result += char
		} else {
			result += "_"
		}
	}
	fmt.Println(result) //Rog_rio
}

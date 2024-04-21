package main

import (
	"fmt"
)

func main() {
	fnfor()
}

func fnfor() {
	i := 1
	for i <= 3 { // 1 2 3
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // 0 1 2
		fmt.Println(j)
	}

	for i := range 3 { // range 0 range 1 range 2
		fmt.Println("range", i)
	}

	for { //loop
		fmt.Println("loop")
		break
	}

	for n := range 6 { //1 3 5
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
// int, float, bool to string
// string to int, float, bool
// int to float, float to int

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Itoa(1), strconv.FormatInt(1, 10)) //1 1 (Equivalent functions)
	fmt.Println(strconv.FormatFloat(12.3456, 'f', 2, 64), strconv.FormatBool(true)) //12.35 true
	fmt.Println(strconv.Atoi("1")) //1 <nil>
	fmt.Println(strconv.ParseFloat("12.3456",64) )//12.3456 <nil>
	fmt.Println(strconv.ParseBool("true")) //true <nil>
	fmt.Println(float64(1)) // 1
	//fmt.Println(int(-1.999)) //error cannot convert 1.999 (untyped float constant) to type int
	x := 1.999
	fmt.Println(int(x)) //1
}


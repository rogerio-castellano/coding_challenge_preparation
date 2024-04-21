package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	l := []int{1, 2, 4, 8, 16}
	l2 := make([]int, len(l))
	copy(l2, l) //dest, src
	//Same behavior as above
	//l2 := append([]int{}, l...)
	fmt.Println(l, len(l), l2) //[1 2 4 8 16] 5 [1 2 4 8 16]
	fmt.Println("Join: ", strings.Join([]string{"a", "b"}, "-")) //Join: a-b
	//fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(l)), ";"), "[]")) // 1;2;4;8;16
	//fmt.Println(strings.Trim(strings.Join(strings.Split(fmt.Sprint(l), " "), ";"), "[]")) // 1;2;4;8;16
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(l), " ", ";", -1), "[]")) // 1;2;4;8;16
	sum := 0
	for _, num := range l {
		sum += num
	}
	fmt.Println(slices.Min(l), slices.Max(l), sum) // 1 16 31
	slices.Sort(l)
	fmt.Println(l[1:3], l[:2], l[1:])  //[2 4] [1 2] [2 4 8 16]
	fmt.Println(slices.Contains(l, 4)) //true
	hasGreaterThan10 := slices.ContainsFunc(l, func(n int) bool {
		return n > 10
	})
	fmt.Println("Has greater than 10:", hasGreaterThan10) //Has greater than 10: true
	fmt.Println("Has negative:", slices.ContainsFunc(l, func(n int) bool {
		return n < 0
	})) // Has negative: false

	fmt.Println(slices.Index(l, 4)) //2
	fmt.Println(slices.IndexFunc(l, func(n int) bool {
		return n > 10
	})) // 4
	
	fmt.Println(l[slices.IndexFunc(l, func(n int) bool {
		return n > 10
	})]) // 16	
}

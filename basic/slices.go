package main

import (
	"fmt"
	"slices"
)

func main() {
	fnslices()
}

func fnslices() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) //uninit: [] true true

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) //emp: [  ] len: 3 cap: 3

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    //set: [a b c]
	fmt.Println("get:", s[2]) //get: c

	fmt.Println("len:", len(s)) //len: 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) //apd: [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) //cpy: [a b c d e f]

	l := s[2:5]
	fmt.Println("sl1:", l) //sl1: [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) //sl2: [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) //sl3: [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) //dcl: [g h i]

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") //t == t2
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) //2d:  [[0] [1 2] [2 3 4]]

	letters := []string{"a", "b", "c", "d", "e"}
	letters = slices.Delete(letters, 1, 4)
	fmt.Println(letters) //[a e]

	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	var index int = 3
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println(numbers) //[10 20 30 50 90 60]

	// creating and getting the capacity of a slice
	sliceOfIntegers := make([]int, 10, 15)
	fmt.Println(cap(sliceOfIntegers)) // 15

	// using append to add multiple values to the slice
	sliceOfIntegers = append(sliceOfIntegers, 5, 6, 7)
	fmt.Println(sliceOfIntegers) // [1 2 3 4 5 6 7]

	// using append to add a slice to a slice
	anotherSlice := []int{8, 9, 10}
	sliceOfIntegers = append(sliceOfIntegers, anotherSlice...)
	fmt.Println(sliceOfIntegers) // [1 2 3 4 5 6 7 8 9 10]

	// creating a map from literals
	studentsAge := map[string]int{
		"solomon": 19,
		"john":    20,
		"janet":   15,
		"daniel":  16,
		"mary":    18,
	}
	fmt.Println(studentsAge) // map[daniel:16 janet:15 john:20 mary:18 solomon:19]

	// deleting keys from the studentsAge map
	delete(studentsAge, "solomon")
	delete(studentsAge, "daniel")

	fmt.Println(studentsAge) // map[janet:15 john:20 mary:18]
}

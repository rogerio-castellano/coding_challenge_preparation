package main

import (
	"fmt"
	"maps"
)

func main() {
	fnmaps()
}

func fnmaps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) //map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1) //v1: 7

	v3 := m["k3"]
	fmt.Println("v3:", v3) //v3: 0

	fmt.Println("len:", len(m)) //len: 2

	delete(m, "k2")
	fmt.Println("map:", m) //map: map[k1:7]

	clear(m)
	fmt.Println("map:", m) //map: map[]

	_, prs := m["k2"]
	fmt.Println("prs:", prs) //prs: false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) //map: map[bar:2 foo:1]

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2") //n == n2
	}

    value, found := n["foo"];
    fmt.Println(value, found) //1 true

	for key, value := range n {
        fmt.Println(key, value) 
    }
}
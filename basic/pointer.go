package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p := person {name : "Rogerio", age : 50}
	commonName(&p)
	changeAge(p)
	x := 1
	increment(&x)
	fmt.Println(x,p)

}

func increment(a *int) {
	*a++
}

func commonName(p *person) {
	p.name = "Joe Smith"
}

func changeAge(p person) {
	p.age = 18
}


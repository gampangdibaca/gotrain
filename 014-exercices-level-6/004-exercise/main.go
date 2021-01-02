package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "I am", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "Bimo",
		last:  "Pangestu",
		age:   21,
	}
	p1.speak()
}

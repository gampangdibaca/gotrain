package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	fmt.Println(p)
	fmt.Println(*p)
	(*p).first = "Henry"
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

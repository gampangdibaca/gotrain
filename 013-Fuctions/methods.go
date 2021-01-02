package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("Hello, my name is", s.first, s.last, "from secretAgent")
}

func (s secretAgent) breath() {
	fmt.Println("Breathing")
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "from Person")
}

type human interface {
	speak()
}

func bars(h human) {
	//type assertion
	switch h.(type) {
	case person:
		fmt.Println("I'm called here,", h.(person).first)
	case secretAgent:
		fmt.Println("I'm called here,", h.(secretAgent).first)
	}
	fmt.Println("called", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()

	p1 := person{
		first: "Dr",
		last:  "No",
	}
	bars(sa1)
	bars(p1)
}

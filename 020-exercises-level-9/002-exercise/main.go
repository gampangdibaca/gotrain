package main

import "fmt"

type person struct {
	First string
	Age   string
}

func (p *person) speak() {
	fmt.Println("Person Speak")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		First: "Budi",
		Age:   "24",
	}
	saySomething(&p1)
	//not work
	//saySomething(p1)
}

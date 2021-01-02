package main

import "fmt"

type person struct {
	first              string
	last               string
	favIceCreamFlavour []string
}

func main() {
	p1 := person{
		first:              "Bimo",
		last:               "Pangestu",
		favIceCreamFlavour: []string{"Chocolate", "Vanilla"},
	}

	p2 := person{
		first:              "Natasha",
		last:               "Olivia",
		favIceCreamFlavour: []string{"Strawberry", "Hazelnut", "Mocca"},
	}

	for _, v := range p1.favIceCreamFlavour {
		fmt.Println(p1.first, p1.last, v)
	}
	for _, v := range p2.favIceCreamFlavour {
		fmt.Println(p2.first, p2.last, v)
	}
}

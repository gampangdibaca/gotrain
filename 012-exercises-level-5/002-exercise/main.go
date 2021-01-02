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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		for _, v2 := range v.favIceCreamFlavour {
			fmt.Println(k, v.first, v.last, v2)
		}
	}

}

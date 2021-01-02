package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//embedded struct
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   25,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
			age:   25,
		},
		ltk: true,
	}
	//sa1.person.first also possible, used when there is identifier name collision
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa2.age, sa2.ltk)

	//anonymous struct
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Josh",
		last:  "Nash",
		age:   22,
	}
	fmt.Println(p3)
}

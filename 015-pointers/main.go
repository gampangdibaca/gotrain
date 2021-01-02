package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a         //operator "&" is to get address
	fmt.Println(*b) //operator "*" is to get the value in that address
	fmt.Println(*&a)
	c := &b
	fmt.Println(**c)
	*b = 43
	fmt.Println(a)
	**c = 44
	fmt.Println(a)
	fmt.Println("---------------------------------")
	x := 0
	foo(x)
	fmt.Println(x)
	fmt.Println("---------------------------------")
	y := 0
	bar(&y)
	fmt.Println(y)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}

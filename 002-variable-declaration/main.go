package main

import "fmt"

var y = 43 //if need to declare outside main body
//DECLARE a VARIABLE with IDENTIFIER "z" with TYPE of int and assigned Zero value
var z int

func main() {
	x := 42 // use as short declaration as much as possible
	fmt.Println(x)

	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("again", y)
}

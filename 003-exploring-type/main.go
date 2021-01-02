package main

import "fmt"

var y = 42
var z = "whatever string"
var a = `Simon said, "whatever string"`
var b int

type hotdog int // my own type

var c hotdog

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// z = 43 will error, because already TYPE string

	//Zero values below this line
	var i string
	var j int
	var k float32
	fmt.Println("print value of i", i, "EOF")
	fmt.Printf("%T\n", i)
	fmt.Println("print value of j", j, "EOF")
	fmt.Printf("%T\n", j)
	fmt.Println("print value of k", k, "EOF")
	fmt.Printf("%T\n", k)
	fmt.Printf("%#V\n", k)

	//Create my own types
	b = 22
	c = 24
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	//conversion vs casting
	// b = c <--- error
	b = int(c)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// there is no casting in Go
}

package main

import "fmt"

func main() {
	foo()
}

func foo() {
	x := 177
	{
		y := 123
		fmt.Println(x + y)
	}
}

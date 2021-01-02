package main

import "fmt"

func main() {
	a := foo()
	a()
}

func foo() func() {
	return func() {
		fmt.Println("Just a value from returned function")
	}
}

package main

import "fmt"

func main() {
	foo(func() int {
		return 123
	})
}

func foo(x func() int) {
	fmt.Println(x())
}

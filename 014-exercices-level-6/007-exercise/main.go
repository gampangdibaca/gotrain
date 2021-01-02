package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("The passed value is", x)
	}
	f(123)
}

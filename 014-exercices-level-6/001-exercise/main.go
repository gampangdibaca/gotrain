package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)

}

func foo() int {
	return 23
}

func bar() (int, string) {
	return 555, "HAHAHA"
}

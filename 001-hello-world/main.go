package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, i am the best in the world ever and ever and ever.")
	foo()
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}
func foo() {
	fmt.Println("I'am in Foo")
}

func bar() {
	x := "wakwakw" //short declaration
	y := 100 + 16
	fmt.Println("exit the control flow", x, y)
}

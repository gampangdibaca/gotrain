package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("First")
	fmt.Println("Second")
	fmt.Println("Third")
	fmt.Println("Last")
}

func foo() {
	fmt.Println("will run at the end")
}

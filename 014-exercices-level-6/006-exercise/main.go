package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("The passed value is", x)
	}(123)
}

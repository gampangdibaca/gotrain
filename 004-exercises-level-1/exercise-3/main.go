package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d, %s, %t\n", x, y, z)
	//s := fmt.Sprintf("%v, %v, %v\n", x, y, z) <--- also working, %v just take the default value
	fmt.Print(s)

}

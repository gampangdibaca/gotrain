package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}

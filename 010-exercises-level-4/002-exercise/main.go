package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5,11,12,13,14,15}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum([]int{2, 3}...))
	fmt.Println("4 + 1 =", mySum([]int{4, 1}...))
	fmt.Println("8 + 6 =", mySum([]int{8, 6}...))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

package main

import (
	"fmt"
	"github.com/gampangdibaca/gotrain/027-testing-and-benchmarking/003-example-test/maroon"
)

func main() {
	fmt.Println("2 + 3 =", maroon.MySum([]int{2, 3}...))
	fmt.Println("4 + 1 =", maroon.MySum([]int{4, 1}...))
	fmt.Println("8 + 6 =", maroon.MySum([]int{8, 6}...))
}

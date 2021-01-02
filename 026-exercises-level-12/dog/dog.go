// Package dog is a package to turn human years to dog years
package dog

import "fmt"

// Years is a func that takes human years of type int and return dog years of type int
func Years(humanYears int) int{
	return humanYears * 7
}

func main() {
	fmt.Println(Years(3))
}
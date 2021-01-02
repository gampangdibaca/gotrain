// Package marron asks if you like maroon color
package maroon

// MySum add an unlimited number of values of type int
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

package main

import "testing"

func TestMySum(t *testing.T) {
	xi := []int{2,3}
	sum := mySum(xi...)
	if sum != 5 {
		t.Error("Expect 5, got", sum)
	}
}

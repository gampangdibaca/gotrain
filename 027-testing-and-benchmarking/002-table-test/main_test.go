package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{
			data:   []int{21, 21},
			answer: 42,
		}, {
			data:   []int{3, 4, 5},
			answer: 12,
		}, {
			data:   []int{1, 1},
			answer: 2,
		}, {
			data:   []int{-1, 2, 5},
			answer: 6,
		},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Error("Expected", v.answer, " got", sum)
		}
	}

}
